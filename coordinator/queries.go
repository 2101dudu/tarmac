package coordinator

import (
	"errors"
	"fmt"
	"strconv"
	"tarmac/cache"
	"tarmac/db"
	"tarmac/json"
	"tarmac/logger"
	"tarmac/utils"
	"tarmac/wsdl"
)

func (s *Service) HandleGetMasterData(url string) (*wsdl.SearchProductMasterDataResponse, error) {
	key := utils.GetSHA256Hash(url)
	collectionName := "master_data"
	id := "master_data"

	data, refreshDB, refreshCache, err := getData(
		s, key, collectionName, id,
		s.wsdlService.SearchProductMasterData,
	)
	if err != nil {
		return nil, err
	}

	go func() {
		if refreshDB {
			s.dbService.RefreshDB(collectionName, id, data)
		}
		if refreshCache {
			s.cacheService.RefreshCache(key, data, s.cacheService.CacheTimes.LongCacheTime)
		}
	}()

	return data, err
}

func (s *Service) HandleSearchProduct(url string) (*wsdl.SearchProductResponse, error) {
	key := utils.GetSHA256Hash(url)
	collectionName := "search_product"
	id := "search_product"

	data, refreshDB, refreshCache, err := getData(
		s, key, collectionName, id,
		s.wsdlService.SearchProducts,
	)

	if err != nil {
		return nil, err
	}

	go func() {
		if refreshDB {
			s.dbService.RefreshDB(collectionName, id, data)
		}
		if refreshCache {
			s.cacheService.RefreshCache(key, data, s.cacheService.CacheTimes.LongCacheTime)
		}
	}()

	return data, nil
}

func (s *Service) HandleSearchProductWithBody(query ProductQuery, url string) ([]*wsdl.Product, *string, *bool, error) {
	key := utils.GetSHA256Hash(url + query.Country + query.Location + query.DepDate + query.Tag)
	collectionName := "search_product_with_body"
	id := key

	resp, refreshDB, err := s.fetchProductsWithFallback(query, key, collectionName, id)
	if err != nil {
		return nil, nil, nil, err
	}

	if err := wsdl.OptimizeProductSearch(resp); err != nil {
		return nil, nil, nil, err
	}

	if refreshDB {
		go func() {
			s.dbService.RefreshDB(collectionName, id, resp)
			s.cacheService.RefreshCache(key, resp, s.cacheService.CacheTimes.LongCacheTime)
		}()
	}

	token := generateToken()
	go s.cacheService.RefreshCache("pagecache:"+token, resp, s.cacheService.CacheTimes.MediumCacheTime)

	firstPage := resp.ProductArray.Items
	if len(firstPage) > query.Length {
		firstPage = firstPage[:query.Length]
	}

	hasMore := len(resp.ProductArray.Items) > query.Length
	return firstPage, &token, &hasMore, nil
}

func (s *Service) fetchProductsWithFallback(query ProductQuery, cacheKey, collectionName, id string) (*wsdl.SearchProductResponse, bool, error) {
	var queriedList []*wsdl.Product
	refreshDB := false

	if cached := cache.CheckCacheHit[wsdl.SearchProductResponse](s.cacheService, cacheKey); cached != nil {
		queriedList = cached.ProductArray.Items
	} else if dbData, _ := db.CheckDBHit[wsdl.SearchProductResponse](s.dbService, collectionName, id); dbData != nil {
		go s.cacheService.RefreshCache(cacheKey, dbData, s.cacheService.CacheTimes.MediumCacheTime)
		queriedList = dbData.ProductArray.Items
	} else {
		refreshDB = true
		productList, dbErr := db.GetAllProducts[wsdl.ProductWrapper](s.dbService, "product_index")
		if dbErr != nil {
			return nil, false, fmt.Errorf("DB fallback failed: %w", dbErr)
		}
		queriedList = wsdl.ApplyQueryToData(productList, query.Country, query.Location, query.DepDate, query.Tag)
	}

	if query.PriceFrom != "" || query.PriceTo != "" || query.NumDays != "" {
		queriedList = wsdl.FilterProducts(queriedList, query.PriceFrom, query.PriceTo, query.NumDays)
	}
	if query.SortBy != "" {
		queriedList = wsdl.SortProducts(queriedList, query.SortBy, query.SortOrder)
	}
	size := len(queriedList)
	totalProducts := strconv.Itoa(size)

	return &wsdl.SearchProductResponse{
		TotalProducts: &totalProducts,
		ProductArray:  &wsdl.ProductArray{Items: queriedList},
	}, refreshDB, nil
}

func (s *Service) HandleSearchProductPagination(token string, cursor, limit int) ([]*wsdl.Product, bool, error) {
	cacheKey := "pagecache:" + token

	fullResp := cache.CheckCacheHit[wsdl.SearchProductResponse](s.cacheService, cacheKey)
	if fullResp == nil || fullResp.ProductArray == nil || fullResp.ProductArray.Items == nil {
		return nil, false, errors.New("pagination cache not found")
	}

	products := fullResp.ProductArray.Items
	if cursor >= len(products) {
		return []*wsdl.Product{}, false, nil
	}

	end := cursor + limit
	if end > len(products) {
		end = len(products)
	}

	return products[cursor:end], end < len(products), nil
}

func (s *Service) HandleDynGetProductParameters(prodCode string) (*wsdl.ProductWithExtraData, error) {
	cacheKey := prodCode
	collectionName := "dyn_product_parameters"
	id := prodCode

	visitedSuccessfully := false
	defer func() {
		if visitedSuccessfully {
			go func() {
				pW, _ := db.CheckDBHit[wsdl.ProductWrapper](s.dbService, "product_index", id)
				if pW != nil {
					pW.VisitCount += 1
					s.dbService.RefreshDB("product_index", id, pW)
				}
			}()
		}
	}()

	if data := cache.CheckCacheHit[wsdl.ProductWithExtraData](s.cacheService, cacheKey); data != nil {
		visitedSuccessfully = true
		return data, nil
	} else if data, _ = db.CheckDBHit[wsdl.ProductWithExtraData](s.dbService, collectionName, id); data != nil {
		go s.cacheService.RefreshCache(cacheKey, data, s.cacheService.CacheTimes.MediumCacheTime)
		visitedSuccessfully = true
		return data, nil
	} else if newData, err := s.wsdlService.DynGetProductParameters(prodCode); err == nil {
		*newData.Name = utils.SimplifyString(*newData.Name)
		productW, _ := db.CheckDBHit[wsdl.ProductWrapper](s.dbService, "product_index", id)

		if productW == nil {
			return nil, errors.New("No product")
		}

		data = &wsdl.ProductWithExtraData{
			Data:             *newData,
			DescriptionArray: extractPointer(productW.Product.TextContentsArray),
			PhotoArray:       extractPointer(productW.Product.PhotoArray),
			Price:            extractPointer(productW.Product.PriceFrom),
		}

		go func() {
			s.dbService.RefreshDB(collectionName, id, data)
			s.cacheService.RefreshCache(cacheKey, data, s.cacheService.CacheTimes.MediumCacheTime)
		}()
		visitedSuccessfully = true
		return data, nil
	} else {
		return nil, err
	}
}

func (s *Service) HandleDynSearchProductAvailableServices(in wsdl.DynProductAvailableServicesRequest) string {
	searchID := generateToken()
	cacheKey := "asyncsearch:" + searchID

	// Save status as "processing"
	s.cacheService.RefreshCache(cacheKey+":status", "processing", s.cacheService.CacheTimes.ShortCacheTime)

	// Spawn goroutine to call SOAP
	go func() {
		resp, err := s.wsdlService.DynSearchProductAvailableServices(in)
		if err != nil {
			logger.Log.Log("Async search failed:", err)
			s.cacheService.RefreshCache(cacheKey+":status", "error", s.cacheService.CacheTimes.ShortCacheTime)
			return
		}

		s.cacheService.RefreshCache(cacheKey+":data", resp, s.cacheService.CacheTimes.MediumCacheTime)
		s.cacheService.RefreshCache(cacheKey+":status", "done", s.cacheService.CacheTimes.MediumCacheTime)
	}()

	return searchID
}

func (s *Service) HandleAsyncAvailableServicesStatus(searchID string) (*string, *wsdl.DynProductAvailableServicesResponse, error) {
	statusKey := "asyncsearch:" + searchID + ":status"
	dataKey := "asyncsearch:" + searchID + ":data"

	status := cache.CheckCacheHit[string](s.cacheService, statusKey)
	if status == nil {
		return nil, nil, errors.New("Search ID not found")
	}

	if *status != "done" {
		return status, nil, nil
	}

	resp := cache.CheckCacheHit[wsdl.DynProductAvailableServicesResponse](s.cacheService, dataKey)
	if resp == nil {
		return nil, nil, errors.New("Corrupted result")
	}
	newStatus := "done"
	return &newStatus, resp, nil
}

type SyncMetadata struct {
	CountryCodes     []string
	LocationCodes    []string
	SeenProductCodes map[string]bool
}

func (s *Service) HandleSyncAllProducts() error {
	var in wsdl.SearchProductRequest
	products, err := s.wsdlService.SearchProductsWithBodyNow(in)
	if err != nil {
		return err
	}

	meta := s.syncProductsToDB(products.ProductArray.Items)
	s.deleteStaleProducts(meta.SeenProductCodes)
	s.deleteEmptyData(meta.CountryCodes, meta.LocationCodes)

	return nil
}

func (s *Service) syncProductsToDB(products []*wsdl.Product) SyncMetadata {
	collectionName := "product_index"

	countrySet := make(map[string]struct{})
	locationSet := make(map[string]struct{})
	seen := make(map[string]bool)

	for _, product := range products {
		if product == nil || product.Code == nil {
			continue
		}
		codeStr := *product.Code
		if _, err := strconv.Atoi(codeStr); err != nil {
			continue
		}

		old, _ := db.CheckDBHit[wsdl.ProductWrapper](s.dbService, collectionName, codeStr)

		tags := []string{}
		enabled := true
		var visitCount int64 = 0
		if old != nil {
			tags = old.Tags
			if old.Enabled != nil {
				enabled = *old.Enabled
			}
			visitCount = old.VisitCount
		}

		wrapped := wsdl.ProductWrapper{
			Product:    *product,
			Tags:       tags,
			Enabled:    &enabled,
			VisitCount: visitCount,
		}
		s.dbService.RefreshDB(collectionName, codeStr, wrapped)

		if product.Country != nil {
			countrySet[*product.Country] = struct{}{}
		}
		if product.Location != nil {
			locationSet[*product.Location] = struct{}{}
		}

		seen[codeStr] = true
	}

	return SyncMetadata{
		CountryCodes:     keys(countrySet),
		LocationCodes:    keys(locationSet),
		SeenProductCodes: seen,
	}
}

func keys(set map[string]struct{}) []string {
	result := make([]string, 0, len(set))
	for k := range set {
		result = append(result, k)
	}
	return result
}

func (s *Service) deleteStaleProducts(seen map[string]bool) {
	products, _ := db.GetAllProducts[wsdl.ProductWrapper](s.dbService, "product_index")

	for _, pw := range products {
		if pw == nil || pw.Product.Code == nil {
			continue
		}
		code := *pw.Product.Code
		if !seen[code] {
			s.dbService.DeleteByID("product_index", code)
			logger.Log.Log("Deleted stale product:", code)
		}
	}
}

func (s *Service) deleteEmptyData(countryCodes, locationCodes []string) {
	collectionName := "master_data"
	id := "master_data"
	unusedKey := ""

	masterData, _, _, err := getData(
		s, unusedKey, collectionName, id,
		s.wsdlService.SearchProductMasterData,
	)

	if err != nil {
		return
	}

	countries := masterData.SearchProductMasterDataArray.CountriesArray.Items
	locations := masterData.SearchProductMasterDataArray.LocationsArray.Items

	newMasterData := &wsdl.SearchProductMasterDataResponse{
		SearchProductMasterDataArray: &wsdl.SearchProductMasterData{
			CountriesArray: &wsdl.CountryArray{Items: []*wsdl.Country{}},
			LocationsArray: &wsdl.LocationArray{Items: []*wsdl.Location{}},
		},
	}

	for _, countryCode := range countryCodes {
		for _, country := range countries {
			if country.Code != nil && *country.Code == countryCode {
				newMasterData.SearchProductMasterDataArray.CountriesArray.Items = append(
					newMasterData.SearchProductMasterDataArray.CountriesArray.Items,
					country,
				)
				break
			}
		}
	}

	for _, locationCode := range locationCodes {
		for _, location := range locations {
			if location.Code != nil && *location.Code == locationCode {
				newMasterData.SearchProductMasterDataArray.LocationsArray.Items = append(
					newMasterData.SearchProductMasterDataArray.LocationsArray.Items,
					location,
				)
				break
			}
		}
	}

	go s.dbService.RefreshDB(collectionName, id, newMasterData)
}

func (s *Service) HandleResetState() {
	s.cacheService.RemoveCache()
	s.dbService.RemoveCollections()
}

func (s *Service) HandleAdminAuth() string {
	token := generateToken()
	go s.cacheService.RefreshCache("admin:"+token, "valid", s.cacheService.CacheTimes.MediumCacheTime)
	return token
}

func (s *Service) HandleCheckAdminAuth(token string) *string {
	return cache.CheckCacheHit[string](s.cacheService, "admin:"+token)
}

func (s *Service) HandleGetAllProducts() ([]*wsdl.ProductWrapper, error) {
	return db.GetAllProducts[wsdl.ProductWrapper](s.dbService, "product_index")
}

func (s *Service) HandleAddProductTags(prodCode string, tags []string) {
	pWrapper, _ := db.CheckDBHit[wsdl.ProductWrapper](s.dbService, "product_index", prodCode)
	pWrapper.Tags = append(pWrapper.Tags, tags...)
	s.dbService.RefreshDB("product_index", prodCode, pWrapper)
}

func (s *Service) HandleRemoveProductTags(tagToRemove, prodCode string) {
	pWrapper, _ := db.CheckDBHit[wsdl.ProductWrapper](s.dbService, "product_index", prodCode)
	// Remove tagToRemove from pWrapper.Tags
	var newTags []string
	for _, tag := range pWrapper.Tags {
		if tag != tagToRemove {
			newTags = append(newTags, tag)
		}
	}
	pWrapper.Tags = newTags
	s.dbService.RefreshDB("product_index", prodCode, pWrapper)
}

func (s *Service) HandleProductToggleState(prodCode string) {
	pW, _ := db.CheckDBHit[wsdl.ProductWrapper](s.dbService, "product_index", prodCode)
	// togle product state
	*pW.Enabled = !*pW.Enabled
	s.dbService.RefreshDB("product_index", prodCode, pW)
}

func (s *Service) HandleHighlightTag(tagName string) {
	s.dbService.RefreshDB("highlighted-tag", "highlighted-tag", tagName)
}

func (s *Service) HandleGetHighlightedTag() (*string, bool) {
	return db.CheckDBHit[string](s.dbService, "highlighted-tag", "highlighted-tag")
}

func (s *Service) HandleDynSetServicesSelectedAndGetToken(in wsdl.DynServicesSelectedRequest) (*string, error) {
	resp, err := s.wsdlService.DynSetServicesSelected(in)
	if err != nil || *resp.Errors.HasErrors != "NO" {
		return nil, err
	}

	b, err := json.Compress(in)
	if err != nil {
		return nil, errors.New("Invalid token")
	}
	token := utils.GetSHA256Hash(string(b))
	cachedSimul := cache.CheckCacheHit[wsdl.DynGetSimulationResponse](s.cacheService, "simulation:"+token)
	if cachedSimul != nil {
		return &token, nil
	}

	var newInput wsdl.DynGetSimulationRequest
	newInput.SessionHash = in.SessionHash
	simul, err := s.wsdlService.DynGetSimulation(newInput)
	if err != nil || simul == nil {
		return nil, errors.New("Invalid Simulation")
	}

	go s.cacheService.RefreshCache("simulation:"+token, simul, s.cacheService.CacheTimes.ShortCacheTime)

	return &token, nil
}

func (s *Service) HandleDynGetSimulation(token string) (*wsdl.DynGetSimulationResponse, error) {
	simul := cache.CheckCacheHit[wsdl.DynGetSimulationResponse](s.cacheService, "simulation:"+token)
	if simul == nil {
		return nil, errors.New("Simulation not found for token—cache may have expired")
	}
	return simul, nil
}
