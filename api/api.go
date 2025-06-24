package api

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"tarmac/cache"
	"tarmac/db"
	"tarmac/env"
	"tarmac/wsdl"
	"time"

	goccyjson "github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gin-contrib/cors"
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

	if refreshDB {
		go db.RefreshDB(a.DBService, collectionName, id, data)
	}

	if refreshCache {
		go cache.RefreshCache(data, key, *a.CacheTimes.LongCacheTime, a.CacheService)
	}

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

	if refreshDB {
		go db.RefreshDB(a.DBService, collectionName, id, data)
	}

	if refreshCache {
		go cache.RefreshCache(data, key, *a.CacheTimes.LongCacheTime, a.CacheService)
	}

	c.JSON(http.StatusOK, data)
}

func (a *Api) handleSearchProductWithBody(c *gin.Context) {
	// Read and preserve request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Parse input
	var input wsdl.SearchProductRequest
	if err := goccyjson.Unmarshal(body, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}
	input.Credentials = a.Credentials
	if input.DepDate != nil {
		// convert strings like "2025-01-09" and "2025-1-9" into
		// the same format, to avoid differnt hash values for the same
		// query
		if canonicalizedDate, err := time.Parse(time.DateOnly, *input.DepDate); err == nil {
			*input.DepDate = canonicalizedDate.String()
		}
	}

	// Key/ID for caching and DB (use full body hash)
	key := getSHA256Hash(c.Request.URL.Path + string(body))
	collectionName := "search_product_with_body"
	id := key // using request-specific key as Mongo ID to make entry unique per request

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

	if refreshDB {
		go func() { // avoid concurrency in mongoDB
			db.RefreshDB(a.DBService, collectionName, id, data)

			for _, product := range data.ProductArray.Items { // refresh every single product entry from the listing
				if product.Code == nil {
					log.Println("Internal ERROR: product code is empty") // temp
					continue
				}

				// check if product is outdated
				existing, outdated := db.CheckDBHit[wsdl.Product](a.DBService, "product_index", *product.Code)
				if existing != nil && !outdated {
					log.Println("Already existing product:", *product.Code) // temp
					continue                                                // fresh, skip
				}
				db.RefreshDB(a.DBService, "product_index", *product.Code, product)
				log.Println("Refreshed product:", *product.Code) // temp
			}
		}()
	}

	if refreshCache {
		go cache.RefreshCache(data, key, *a.CacheTimes.MediumCacheTime, a.CacheService)
	}

	// Page token for pagination
	token := generateToken()
	if jsonData, err := goccyjson.Marshal(data); err == nil {
		a.CacheService.Set("pagecache:"+token, jsonData, *a.CacheTimes.MediumCacheTime)
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
	token := c.Query("token")
	cursorStr := c.Query("cursor")
	limitStr := c.Query("limit")

	if token == "" || cursorStr == "" || limitStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing token, cursor or limit query parameter"})
		return
	}

	cursor, err := strconv.Atoi(cursorStr)
	if err != nil || cursor < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursor"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
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
	prodCode := c.Param("prodCode")
	if _, err := strconv.Atoi(prodCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
		return
	}

	key := getSHA256Hash(c.Request.URL.Path + "?" + c.Request.URL.RawQuery)
	collectionName := "dyn_product_parameters"
	id := prodCode

	data, refreshDB, refreshCache, err := getData(
		a, key, collectionName, id,
		func() (*wsdl.DynProductParametersResponse, error) {
			return a.SoapService.DynGetProductParameters(&wsdl.DynProductParametersRequest{
				Credentials: a.Credentials,
				Productcode: &prodCode,
			})
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if refreshDB {
		go db.RefreshDB(a.DBService, collectionName, id, data)
	}

	if refreshCache {
		go cache.RefreshCache(data, key, *a.CacheTimes.MediumCacheTime, a.CacheService)
	}

	c.JSON(http.StatusOK, data)
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
			log.Println()
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
