package api

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"tarmac/api/cache"
	"tarmac/api/helper"
	"tarmac/env"
	"tarmac/internal/render"
	"tarmac/wsdl"
	"time"

	goccyjson "github.com/goccy/go-json"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Api struct {
	SoapService wsdl.Wbs_pkt_methodsSoap
	Credentials *wsdl.CredentialsStruct
	CacheTimes  *env.CacheTimes
	DBService   *redis.Client
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
	cacheHandler(c, getSHA256Hash(c.Request.URL.String()), *a.CacheTimes.LongCacheTime, a.DBService, func() (*wsdl.SearchProductMasterDataResponse, error) {
		return a.SoapService.SearchProductMasterData(&wsdl.SearchProductMasterDataRequest{
			Credentials: a.Credentials,
		})
	})
}

func (a *Api) handleSearchProduct(c *gin.Context) {
	cacheHandler(c, getSHA256Hash(c.Request.URL.String()), *a.CacheTimes.LongCacheTime, a.DBService, func() (*wsdl.SearchProductResponse, error) {
		return a.SoapService.SearchProducts(&wsdl.SearchProductRequest{
			Credentials: a.Credentials,
		})
	})
}

func (a *Api) handleSearchProductWithBody(c *gin.Context) {
	// preserve request body (BindJSON consumes it)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	cacheKey := getSHA256Hash(c.Request.URL.Path + string(body))

	cached := cache.CheckCacheHit[wsdl.SearchProductResponse](cacheKey, a.DBService)
	if cached != nil { // cache hit
		token := generateToken()
		if jsonData, err := goccyjson.Marshal(cached); err == nil {
			a.DBService.Set("pagecache:"+token, jsonData, *a.CacheTimes.MediumCacheTime)
		}
		firstPage := cached.ProductArray.Items
		if len(firstPage) > 24 {
			firstPage = firstPage[:24]
		}
		c.JSON(http.StatusOK, gin.H{
			"products": firstPage,
			"token":    token,
			"hasMore":  len(cached.ProductArray.Items) > 24,
		})
		return
	}

	// Not cached — parse input
	var input wsdl.SearchProductRequest
	if err := goccyjson.Unmarshal(body, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}
	input.Credentials = a.Credentials

	// SOAP call
	response, err := a.SoapService.SearchProducts(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products: " + err.Error()})
		return
	}
	if err := helper.OptimizeProductSearch(response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid number of products: " + err.Error()})
		return
	}

	// Marshal + set in cache
	jsonData, err := goccyjson.Marshal(response)
	if err == nil {
		a.DBService.Set(cacheKey, jsonData, *a.CacheTimes.MediumCacheTime)
	}

	// Generate token and store full result separately
	token := generateToken()
	if err == nil {
		a.DBService.Set("pagecache:"+token, jsonData, *a.CacheTimes.MediumCacheTime)
	}

	// Return only first 24 products
	firstPage := response.ProductArray.Items
	if len(firstPage) > 24 {
		firstPage = firstPage[:24]
	}

	c.JSON(http.StatusOK, gin.H{
		"products": firstPage,
		"token":    token,
		"hasMore":  len(response.ProductArray.Items) > 24,
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cursor must be an integer"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit must be an integer"})
		return
	}

	// Get full list from Redis
	val, err := a.DBService.Get("pagecache:" + token).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found or expired"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var fullResp *wsdl.SearchProductResponse
	if err := goccyjson.Unmarshal([]byte(val), &fullResp); err != nil {
		a.DBService.Del("pagecache:" + token)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Corrupted cache, cleared"})
		return
	}

	// Pagination
	products := fullResp.ProductArray
	if products == nil || cursor >= len(products.Items) {
		c.JSON(http.StatusOK, gin.H{"products": []wsdl.Product{}})
		return
	}

	end := cursor + limit
	if end > len(products.Items) {
		end = len(products.Items)
	}

	hasMore := end < len(products.Items)

	paged := products.Items[cursor:end]
	c.JSON(http.StatusOK, gin.H{"products": paged, "hasMore": hasMore})
}

func (a *Api) handleDynGetProductParameters(c *gin.Context) {
	prodCode := c.Param("prodCode")
	if _, err := strconv.Atoi(prodCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
		return
	}

	cacheHandler[wsdl.DynProductParametersResponse](
		c,
		getSHA256Hash(c.Request.URL.Path+"?"+c.Request.URL.RawQuery),
		*a.CacheTimes.MediumCacheTime,
		a.DBService,
		func() (*wsdl.DynProductParametersResponse, error) {
			return a.SoapService.DynGetProductParameters(&wsdl.DynProductParametersRequest{
				Credentials: a.Credentials,
				Productcode: &prodCode,
			})
		},
	)
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

func cacheHandler[T any](c *gin.Context, key string, ttl time.Duration, db *redis.Client, fetch func() (*T, error)) {
	dataInCache := cache.CheckCacheHit[T](key, db)
	if dataInCache != nil {
		c.Render(http.StatusOK, render.JSON{Data: dataInCache})
		return
	}

	resp, err := fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := goccyjson.Marshal(resp)
	if err == nil {
		db.Set(key, jsonData, ttl)
	}

	c.Render(http.StatusOK, render.JSON{Data: resp})
}
