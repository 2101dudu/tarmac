package api

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"tarmac/api/cronjob"
	"tarmac/api/helper"
	"tarmac/cache"
	"tarmac/db"
	"tarmac/env"
	"tarmac/logger"
	"tarmac/mail"
	"tarmac/wsdl"

	goccyjson "github.com/goccy/go-json"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"
)

type Api struct {
	SoapService         wsdl.Wbs_pkt_methodsSoap
	Credentials         *wsdl.CredentialsStruct
	CacheTimes          *env.CacheTimes
	CacheService        *redis.Client
	DBService           *db.Service
	MailService         *mail.MailService
	AdminHashedPassword string
}

func (a *Api) Start() {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.1.120:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	engine.GET("api/get/master-data", a.handleGetMasterData)
	engine.GET("api/search/product", a.handleSearchProduct)
	engine.POST("api/search/product", a.handleSearchProductWithBody)
	engine.GET("api/search/product/page", a.handleSearchProductPagination)
	engine.GET("api/get/product/:prodCode", a.handleDynGetProductParameters)
	engine.POST("api/dynamic/product/available-services", a.handleDynSearchProductAvailableServices)
	engine.GET("api/dynamic/product/available-services/status", a.handleAsyncAvailableServicesStatus)
	engine.POST("api/dynamic/product/set-services", a.handleDynSetServicesSelected)
	engine.POST("api/dynamic/product/check-set-services", a.handleDynCheckSetServices)
	engine.POST("api/dynamic/product/get-simulation", a.handleDynGetSimulation)
	engine.GET("api/page/highlighted/tag", a.handleGetHighlightedTag)

	//========================================================================
	engine.POST("api/admin/auth", a.handleAdminAuth) // ADMIN ENDPOINT

	admin := engine.Group("/api/admin")
	admin.Use(a.AdminAuthMiddleware())
	admin.GET("/products", a.handleListAllProducts)
	admin.POST("/products/:prodCode/tags/add", a.handleAddProductTags)
	admin.GET("/products/:prodCode/tags/remove/:tagToRemove", a.handleRemoveProductTags)
	admin.POST("/products/:prodCode/disable", a.handleProductDisable)
	admin.POST("/products/:prodCode/enable", a.handleProductEnable)
	admin.GET("/page/highlight/:tagToHighlight", a.handleHighlightTag)

	//========================================================================

	//========================================================================
	engine.GET("api/admin/reset-state", a.handleResetState) // DEV PURPOSE
	//========================================================================

	//========================================================================
	engine.GET("api/admin/sync/products", a.handleSyncAllProducts) // MANUAL CALL

	c := cron.New()
	c.AddFunc("0 */6 * * *", func() {
		_, cc, lc := cronjob.SyncAllProducts(a.SoapService, a.Credentials, a.DBService)
		a.helperFuncSyncAllProducts(cc, lc)
	})
	c.Start()
	//========================================================================

	engine.Run(":8080")
}

func (a *Api) handleGetMasterData(c *gin.Context) {
	key := getSHA256Hash(c.Request.URL.String())
	collectionName := "master_data"
	id := "master_data"

	data, refreshDB, refreshCache, err := getData(
		a, key, collectionName, id,
		func() (*wsdl.SearchProductMasterDataResponse, error) {
			return a.SoapService.SearchProductMasterData(&wsdl.SearchProductMasterDataRequest{
				Credentials: a.Credentials,
			})
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	maybeRefreshDBAndCache(refreshDB, refreshCache,
		func() { a.DBService.RefreshDB(collectionName, id, data) },
		func() { cache.RefreshCache(data, key, *a.CacheTimes.LongCacheTime, a.CacheService) },
	)

	c.JSON(http.StatusOK, data)
}

func (a *Api) handleSearchProduct(c *gin.Context) {
	key := getSHA256Hash(c.Request.URL.String())
	collectionName := "search_product"
	id := "search_product"

	data, refreshDB, refreshCache, err := getData(
		a, key, collectionName, id,
		func() (*wsdl.SearchProductResponse, error) {
			return a.SoapService.SearchProducts(&wsdl.SearchProductRequest{
				Credentials: a.Credentials,
			})
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	maybeRefreshDBAndCache(refreshDB, refreshCache,
		func() { a.DBService.RefreshDB(collectionName, id, data) },
		func() { cache.RefreshCache(data, key, *a.CacheTimes.LongCacheTime, a.CacheService) },
	)

	c.JSON(http.StatusOK, data)
}

func (a *Api) handleSearchProductWithBody(c *gin.Context) {
	// Read the full body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Unmarshal to generic map
	var raw map[string]any
	if err := goccyjson.Unmarshal(body, &raw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	// list length
	listLength := 24 // default 24 pages
	if listLengthFromQuery, err := strconv.Atoi(extractStringField(raw, "Length")); err == nil && listLengthFromQuery > 0 {
		listLength = listLengthFromQuery
	}

	// sort
	country := extractStringField(raw, "Country")
	location := extractStringField(raw, "Location")
	depDate := extractStringField(raw, "DepDate")
	tagName := extractStringField(raw, "Tag")

	// for the best-sellers page
	if tagName == "highlighted-tag" {
		tagNameOnDB, _ := db.CheckDBHit[string](a.DBService, "highlighted-tag", "highlighted-tag")
		if tagNameOnDB != nil {
			tagName = *tagNameOnDB
		}
	}

	sortBy := extractStringField(raw, "SortBy")
	sortOrder := extractStringField(raw, "SortOrder")

	// filter
	priceFrom := extractStringField(raw, "PriceFrom")
	priceTo := extractStringField(raw, "PriceTo")
	days := extractStringField(raw, "NumDays")

	// Hash using re-marshaled cleanBody
	key := getSHA256Hash(c.Request.URL.Path + country + location + depDate + tagName)
	collectionName := "search_product_with_body"
	id := key

	var queriedList []*wsdl.Product
	refreshDB := false

	// Try cache first
	if cached := cache.CheckCacheHit[wsdl.SearchProductResponse](key, a.CacheService); cached != nil {
		queriedList = cached.ProductArray.Items
	} else if dbData, _ := db.CheckDBHit[wsdl.SearchProductResponse](a.DBService, collectionName, id); dbData != nil {
		// Fallback to DB if not in cache
		go cache.RefreshCache(dbData, key, *a.CacheTimes.MediumCacheTime, a.CacheService)
		queriedList = dbData.ProductArray.Items
	} else {
		// Fallback to full DB search if not in cache or DB
		refreshDB = true
		logger.Log.Log("Falling back to full DB search")
		productList, dbErr := db.GetAllProducts[wsdl.ProductWrapper](a.DBService, "product_index")
		if dbErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB fallback failed: " + dbErr.Error()})
			return
		}
		queriedList = helper.ApplyQueryToData(productList, country, location, depDate, tagName)
	}

	if priceFrom != "" || priceTo != "" || days != "" {
		queriedList = helper.FilterProducts(queriedList, priceFrom, priceTo, days)
	}
	if sortBy != "" {
		queriedList = helper.SortProducts(queriedList, sortBy, sortOrder)
	}
	size := len(queriedList)
	totalProducts := strconv.Itoa(size)

	response := &wsdl.SearchProductResponse{
		TotalProducts: &totalProducts,
		ProductArray:  &wsdl.ProductArray{Items: queriedList},
	}

	err = helper.OptimizeProductSearch(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to optimize product search: " + err.Error()})
		return
	}

	if refreshDB {
		go func() {
			a.DBService.RefreshDB(collectionName, id, response)
			cache.RefreshCache(response, key, *a.CacheTimes.LongCacheTime, a.CacheService)
		}()
	}

	// Page token for pagination
	token := generateToken()
	go func() {
		if jsonData, err := goccyjson.Marshal(response); err == nil {
			err := a.CacheService.Set("pagecache:"+token, jsonData, *a.CacheTimes.MediumCacheTime).Err()
			if err != nil {
				logger.Log.Log("Failed to set cache page:", err)
			}
		}
	}()

	// Slice to first listLength
	firstPage := response.ProductArray.Items
	if len(firstPage) > listLength {
		firstPage = firstPage[:listLength]
	}

	c.JSON(http.StatusOK, gin.H{
		"products": firstPage,
		"token":    token,
		"hasMore":  len(response.ProductArray.Items) > listLength,
	})
}

func (a *Api) handleSearchProductPagination(c *gin.Context) {
	token, cursor, limit, err := parsePaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cacheKey := "pagecache:" + token

	// Retrieve full response from cache
	val, err := a.CacheService.Get(cacheKey).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found or expired"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cache error: " + err.Error()})
		return
	}

	var fullResp *wsdl.SearchProductResponse
	if err := goccyjson.Unmarshal([]byte(val), &fullResp); err != nil || fullResp == nil {
		_ = a.CacheService.Del(cacheKey) // clear corrupted entry
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Corrupted cache. Entry cleared."})
		return
	}

	products := fullResp.ProductArray.Items
	if cursor >= len(products) {
		c.JSON(http.StatusOK, gin.H{"products": []wsdl.Product{}, "hasMore": false})
		return
	}

	end := cursor + limit
	if end > len(products) {
		end = len(products)
	}

	paged := products[cursor:end]
	hasMore := end < len(products)

	c.JSON(http.StatusOK, gin.H{
		"products": paged,
		"hasMore":  hasMore,
	})
}

func (a *Api) handleDynGetProductParameters(c *gin.Context) {
	prodCodeRaw := c.Param("prodCode")
	_, err := strconv.Atoi(prodCodeRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
		return
	}

	cacheKey := prodCodeRaw
	collectionName := "dyn_product_parameters"
	id := prodCodeRaw

	type productWithExtradData struct {
		Data             wsdl.DynProductParametersResponse `json:"data"`
		DescriptionArray wsdl.TextContentArray             `json:"descriptionArray"`
		PhotoArray       wsdl.PhotoArray                   `json:"photoArray"`
		Price            string                            `json:"price"`
	}

	// Check cache first
	cachedData := cache.CheckCacheHit[productWithExtradData](cacheKey, a.CacheService)
	if cachedData != nil {
		c.JSON(http.StatusOK, *cachedData)
		return
	}

	// Check DB next
	dbData, refreshDB := db.CheckDBHit[productWithExtradData](a.DBService, collectionName, id)
	if dbData != nil {
		go cache.RefreshCache(dbData, cacheKey, *a.CacheTimes.MediumCacheTime, a.CacheService)
		c.JSON(http.StatusOK, *dbData)
		if !refreshDB {
			return // fresh, skip fetching from API
		}
	}

	// Not in cache or DB, fetch data from API
	data, err := a.SoapService.DynGetProductParameters(&wsdl.DynProductParametersRequest{
		Credentials: a.Credentials,
		Productcode: &prodCodeRaw,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	*data.Name = helper.SimplifyString(*data.Name)

	productW, _ := db.CheckDBHit[wsdl.ProductWrapper](a.DBService, "product_index", id)

	if productW == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No product"})
		logger.Log.Log("[WARNING] No product:", prodCodeRaw)
		return
	}

	product := productW.Product

	fullData := productWithExtradData{
		Data:             *data,
		DescriptionArray: extractPointer(product.TextContentsArray),
		PhotoArray:       extractPointer(product.PhotoArray),
		Price:            extractPointer(product.PriceFrom),
	}

	go a.DBService.RefreshDB(collectionName, id, fullData)
	go cache.RefreshCache(fullData, cacheKey, *a.CacheTimes.MediumCacheTime, a.CacheService)

	c.JSON(http.StatusOK, fullData)
}

func (a *Api) handleDynSearchProductAvailableServices(c *gin.Context) {
	var input wsdl.DynProductAvailableServicesRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	input.Credentials = a.Credentials
	searchId := generateToken()
	cacheKey := "asyncsearch:" + searchId

	// Save status as "processing"
	cache.RefreshCache("processing", cacheKey+":status", *a.CacheTimes.ShortCacheTime, a.CacheService)

	// Spawn goroutine to call SOAP
	go func() {
		resp, err := a.SoapService.DynSearchProductAvailableServices(&input)
		if err != nil {
			logger.Log.Log("Async search failed:", err)
			cache.RefreshCache("error", cacheKey+":status", *a.CacheTimes.ShortCacheTime, a.CacheService)
			return
		}

		cache.RefreshCache(resp, cacheKey+":data", *a.CacheTimes.MediumCacheTime, a.CacheService)
		cache.RefreshCache("done", cacheKey+":status", *a.CacheTimes.MediumCacheTime, a.CacheService)
	}()

	// Return searchId immediately
	c.JSON(http.StatusOK, gin.H{"searchId": searchId})
}

func (a *Api) handleAsyncAvailableServicesStatus(c *gin.Context) {
	searchId := c.Query("id")
	if searchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing search ID"})
		return
	}

	statusKey := "asyncsearch:" + searchId + ":status"
	dataKey := "asyncsearch:" + searchId + ":data"

	status := cache.CheckCacheHit[string](statusKey, a.CacheService)
	if status == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Search ID not found"})
		return
	}

	if *status != "done" {
		c.JSON(http.StatusOK, gin.H{"status": status})
		return
	}

	resp := cache.CheckCacheHit[wsdl.DynProductAvailableServicesResponse](dataKey, a.CacheService)
	if resp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Corrupted result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "done",
		"data":   resp,
	})
}

func (a *Api) handleDynSetServicesSelected(c *gin.Context) {
	var input wsdl.DynServicesSelectedRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	input.Credentials = a.Credentials
	resp, err := a.SoapService.DynSetServicesSelected(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Render(http.StatusOK, render.JSON{Data: resp})
}

func (a *Api) handleDynCheckSetServices(c *gin.Context) {
	var input wsdl.DynGetOptionalsSelectedRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	input.Credentials = a.Credentials
	resp, err := a.SoapService.DynGetOptionalsSelected(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Render(http.StatusOK, render.JSON{Data: resp})
}

func (a *Api) handleDynGetSimulation(c *gin.Context) {
	var input wsdl.DynGetSimulationRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	input.Credentials = a.Credentials
	resp, err := a.SoapService.DynGetSimulation(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Render(http.StatusOK, render.JSON{Data: resp})
}

// just handles cache + DB + fetch. returns (data, refreshDB, refreshCache, error)
func getData[T any](a *Api, cacheKey, collectionName, id string, fetchFunc func() (T, error)) (T, bool, bool, error) {
	// defer logger.Log.TrackTime()()
	var empty T

	// Check cache first
	cachedData := cache.CheckCacheHit[T](cacheKey, a.CacheService)
	if cachedData != nil {
		return *cachedData, false, false, nil
	}

	// Check DB next
	dbData, outdated := db.CheckDBHit[T](a.DBService, collectionName, id)
	if dbData != nil {
		if outdated {
			logger.Log.Log("Outdated DB data")
		}
		return *dbData, outdated, true, nil
	}

	// Not in cache or DB, fetch from API
	newData, err := fetchFunc()
	if err != nil {
		return empty, true, true, err
	}

	return newData, true, true, nil
}

func maybeRefreshDBAndCache(refreshDB, refreshCache bool, dbFunc func(), cacheFunc func()) {
	if refreshDB {
		go dbFunc()
	}
	if refreshCache {
		go cacheFunc()
	}
}

func parsePaginationParams(c *gin.Context) (token string, cursor, limit int, err error) {
	token = c.Query("token")
	cursorStr := c.Query("cursor")
	limitStr := c.Query("limit")

	if token == "" || cursorStr == "" || limitStr == "" {
		err = errors.New("Missing token, cursor or limit query parameter")
		return
	}

	cursor, err = strconv.Atoi(cursorStr)
	if err != nil || cursor < 0 {
		err = errors.New("Invalid cursor")
		return
	}

	limit, err = strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		err = errors.New("Invalid limit")
		return
	}

	return
}

func (a *Api) handleResetState(c *gin.Context) {
	go func() {
		if err := a.CacheService.FlushAll().Err(); err != nil {
			logger.Log.Log("Failed to flush Redis:", err)
		} else {
			logger.Log.Log("Redis cache cleared")
		}
	}()
	go a.DBService.RemoveCollections()

	c.JSON(http.StatusOK, gin.H{"status": "reset concluded"})
}

func (a *Api) handleSyncAllProducts(c *gin.Context) {
	err, countryCodes, locationCodes := cronjob.SyncAllProducts(a.SoapService, a.Credentials, a.DBService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync: " + err.Error()})
		return
	}
	a.helperFuncSyncAllProducts(countryCodes, locationCodes)

	c.JSON(http.StatusOK, gin.H{"message": "Product sync ended."})
}

func (a *Api) helperFuncSyncAllProducts(countryCodes, locationCodes []string) {
	collectionName := "master_data"
	id := "master_data"
	unusedKey := ""

	masterData, _, _, err := getData(
		a, unusedKey, collectionName, id,
		func() (*wsdl.SearchProductMasterDataResponse, error) {
			return a.SoapService.SearchProductMasterData(&wsdl.SearchProductMasterDataRequest{
				Credentials: a.Credentials,
			})
		},
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

	go a.DBService.RefreshDB(collectionName, id, newMasterData)
}

func (a *Api) handleAdminAuth(c *gin.Context) {
	var body struct{ Password string }
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	expected := a.AdminHashedPassword
	if getSHA256Hash(body.Password) != expected {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token := generateToken()
	go cache.RefreshCache("valid", "admin:"+token, *a.CacheTimes.MediumCacheTime, a.CacheService)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *Api) AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		if storedToken := cache.CheckCacheHit[string]("admin:"+token, a.CacheService); *storedToken != "valid" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func (a *Api) handleListAllProducts(c *gin.Context) {
	productList, dbErr := db.GetAllProducts[wsdl.ProductWrapper](a.DBService, "product_index")
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB failed: " + dbErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": productList})
}

func (a *Api) handleAddProductTags(c *gin.Context) {
	prodCodeRaw := c.Param("prodCode")
	_, err := strconv.Atoi(prodCodeRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
		return
	}

	var req struct {
		Tags []string `json:"tags"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	go func() {
		pWrapper, _ := db.CheckDBHit[wsdl.ProductWrapper](a.DBService, "product_index", prodCodeRaw)
		pWrapper.Tags = append(pWrapper.Tags, req.Tags...)
		a.DBService.RefreshDB("product_index", prodCodeRaw, pWrapper)
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Tags added successfully"})
}

func (a *Api) handleRemoveProductTags(c *gin.Context) {
	prodCodeRaw := c.Param("prodCode")
	_, err := strconv.Atoi(prodCodeRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
		return
	}

	tagToRemove := c.Param("tagToRemove")

	if tagToRemove == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing option to remove"})
		return
	}

	go func() {
		pWrapper, _ := db.CheckDBHit[wsdl.ProductWrapper](a.DBService, "product_index", prodCodeRaw)
		// Remove tagToRemove from pWrapper.Tags
		var newTags []string
		for _, tag := range pWrapper.Tags {
			if tag != tagToRemove {
				newTags = append(newTags, tag)
			}
		}
		pWrapper.Tags = newTags
		a.DBService.RefreshDB("product_index", prodCodeRaw, pWrapper)
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Tag removed successfully"})
}

func (a *Api) handleProductDisable(c *gin.Context) {
	prodCodeRaw := c.Param("prodCode")
	_, err := strconv.Atoi(prodCodeRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
		return
	}

	pW, _ := db.CheckDBHit[wsdl.ProductWrapper](a.DBService, "product_index", prodCodeRaw)

	// disable product
	*pW.Enabled = false

	a.DBService.RefreshDB("product_index", prodCodeRaw, pW)

	c.JSON(http.StatusOK, gin.H{"message": "Product disabled successfully"})
}

func (a *Api) handleProductEnable(c *gin.Context) {
	prodCodeRaw := c.Param("prodCode")
	_, err := strconv.Atoi(prodCodeRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
		return
	}

	pW, _ := db.CheckDBHit[wsdl.ProductWrapper](a.DBService, "product_index", prodCodeRaw)

	// disable product
	*pW.Enabled = true

	a.DBService.RefreshDB("product_index", prodCodeRaw, pW)

	c.JSON(http.StatusOK, gin.H{"message": "Product enabled successfully"})
}

func (a *Api) handleHighlightTag(c *gin.Context) {
	tagName := c.Param("tagToHighlight")
	a.DBService.RefreshDB("highlighted-tag", "highlighted-tag", tagName)
	c.JSON(http.StatusOK, gin.H{"message": "Product enabled successfully"})
}

func (a *Api) handleGetHighlightedTag(c *gin.Context) {
	tagName, _ := db.CheckDBHit[string](a.DBService, "highlighted-tag", "highlighted-tag")
	if tagName == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No highlighted tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tag": *tagName})
}
