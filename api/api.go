package api

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
	"tarmac/api/helper"
	"tarmac/cache"
	"tarmac/db"
	"tarmac/env"
	"tarmac/logger"
	"tarmac/wsdl"
	"time"

	goccyjson "github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/go-redis/redis"
)

type Api struct {
	SoapService  wsdl.Wbs_pkt_methodsSoap
	Credentials  *wsdl.CredentialsStruct
	CacheTimes   *env.CacheTimes
	CacheService *redis.Client
	DBService    *mongo.Client
}

func (a *Api) Start() {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.1.120:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	engine.GET("api/get/master-data", a.handleGetMasterData)
	engine.GET("api/search/product", a.handleSearchProduct)
	engine.POST("api/search/product", a.handleSearchProductWithBody)
	engine.GET("api/search/product/page", a.handleSearchProductPagination)
	engine.GET("api/get/product/:prodCode", a.handleDynGetProductParameters)
	engine.POST("api/dynamic/product/available-services", a.handleDynSearchProductAvailableServices)
	engine.POST("api/dynamic/product/set-services", a.handleDynSetServicesSelected)
	engine.POST("api/dynamic/product/check-set-services", a.handleDynCheckSetServices)
	engine.POST("api/dynamic/product/get-simulation", a.handleDynGetSimulation)

	//========================================================================
	engine.POST("api/admin/reset-state", a.handleResetState) // DEV PURPOSE
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
		func() { db.RefreshDB(a.DBService, collectionName, id, data) },
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
		func() { db.RefreshDB(a.DBService, collectionName, id, data) },
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

	// sort
	sortBy := extractStringField(raw, "SortBy")
	sortOrder := extractStringField(raw, "SortOrder")

	// filter
	priceFrom := extractStringField(raw, "PriceFrom")
	priceTo := extractStringField(raw, "PriceTo")
	days := extractStringField(raw, "NumDays")

	// Marshal back only the relevant search fields for the WSDL
	cleanBody, err := goccyjson.Marshal(raw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process input"})
		return
	}

	// Now unmarshal into your actual WSDL struct
	var input wsdl.SearchProductRequest
	if err := goccyjson.Unmarshal(cleanBody, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input for WSDL: " + err.Error()})
		return
	}
	input.Credentials = a.Credentials

	// Normalize dep date
	if input.DepDate != nil {
		// convert strings like "2025-01-09" and "2025-1-9" into
		// the same format, to avoid differnt hash values for the same
		// query
		if canonicalizedDate, err := time.Parse(time.DateOnly, *input.DepDate); err == nil {
			*input.DepDate = canonicalizedDate.String()
		}
	}

	// Hash using re-marshaled cleanBody
	key := getSHA256Hash(c.Request.URL.Path + string(cleanBody))
	collectionName := "search_product_with_body"
	id := key

	// Cache + DB dance
	data, refreshDB, refreshCache, err := getData(
		a, key, collectionName, id,
		func() (*wsdl.SearchProductResponse, error) {
			return a.SoapService.SearchProducts(&input)
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = helper.OptimizeProductSearch(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to optimize product search: " + err.Error()})
		return
	}

	if refreshDB {
		go func() { // avoid concurrency in mongoDB
			db.RefreshDB(a.DBService, collectionName, id, data)

			for _, product := range data.ProductArray.Items { // refresh every single product entry from the listing
				if product.Code == nil {
					logger.Log.Log("Internal ERROR: product code is empty") // temp
					continue
				}
				prodCodeInt, err := strconv.Atoi(*product.Code)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
					return
				}
				prodCode := strconv.Itoa(prodCodeInt)

				// check if product is outdated
				existing, outdated := db.CheckDBHit[wsdl.Product](a.DBService, "product_index", prodCode)
				if existing != nil && !outdated {
					continue // fresh, skip
				}
				db.RefreshDB(a.DBService, "product_index", prodCode, product)
			}
		}()
	}

	if refreshCache {
		go cache.RefreshCache(data, key, *a.CacheTimes.MediumCacheTime, a.CacheService)
	}

	if sortBy != "" {
		data.ProductArray.Items = helper.SortProducts(data.ProductArray.Items, sortBy, sortOrder)
	}

	if priceFrom != "" || priceTo != "" || days != "" {
		data.ProductArray.Items = helper.FilterProducts(data.ProductArray.Items, priceFrom, priceTo, days)
	}

	// Page token for pagination
	token := generateToken()
	if jsonData, err := goccyjson.Marshal(data); err == nil {
		go func() {
			err := a.CacheService.Set("pagecache:"+token, jsonData, *a.CacheTimes.MediumCacheTime).Err()
			if err != nil {
				logger.Log.Log("Failed to set cache page:", err)
			}
		}()
	}

	// Slice to first 24
	firstPage := data.ProductArray.Items
	if len(firstPage) > 24 {
		firstPage = firstPage[:24]
	}

	c.JSON(http.StatusOK, gin.H{
		"products": firstPage,
		"token":    token,
		"hasMore":  len(data.ProductArray.Items) > 24,
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
	prodCodeInt, err := strconv.Atoi(prodCodeRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
		return
	}
	prodCode := strconv.Itoa(prodCodeInt)

	cacheKey := prodCode
	collectionName := "dyn_product_parameters"
	id := prodCode

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
		Productcode: &prodCode,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: check if "outdated" matters...
	// ...., outdated := ...
	product, _ := db.CheckDBHit[wsdl.Product](a.DBService, "product_index", id)

	if product == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No product"})
		logger.Log.Log("[WARNING] No product:", prodCode)
		return
	}

	fullData := productWithExtradData{
		Data:             *data,
		DescriptionArray: extractPointer(product.TextContentsArray),
		PhotoArray:       extractPointer(product.PhotoArray),
		Price:            extractPointer(product.PriceFrom),
	}

	go db.RefreshDB(a.DBService, collectionName, id, fullData)
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
	resp, err := a.SoapService.DynSearchProductAvailableServices(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Render(http.StatusOK, render.JSON{Data: resp})
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
	defer logger.Log.TrackTime()()
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
		// Flush Redis
		if err := a.CacheService.FlushAll().Err(); err != nil {
			logger.Log.Log("Failed to flush Redis:", err)
		} else {
			logger.Log.Log("Redis cache cleared")
		}
	}()

	go func() {
		// Drop all MongoDB collections in your DB (be careful)
		collections, err := a.DBService.Database("tarmac").ListCollectionNames(c, struct{}{})
		if err != nil {
			logger.Log.Log("Failed to list Mongo collections:", err)
			return
		}
		for _, col := range collections {
			if err := a.DBService.Database("tarmac").Collection(col).Drop(c); err != nil {
				logger.Log.Log("Failed to drop collection "+col+":", err)
			} else {
				logger.Log.Log("Dropped collection:", col)
			}
		}
	}()

	c.JSON(http.StatusOK, gin.H{"status": "reset initiated"})
}
