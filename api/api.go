package api

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"tarmac/api/cache"
	"tarmac/api/helper"
	"tarmac/cache"
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
	cacheHandlerWithBody(
		c,
		*a.CacheTimes.MediumCacheTime,
		a.DBService,
		func(req *wsdl.SearchProductRequest) (*wsdl.SearchProductResponse, error) {
			req.Credentials = a.Credentials
			resp, err := a.SoapService.SearchProducts(req)
			if err != nil {
				return nil, err
			}
			return resp, helper.OptimizeProductSearch(resp)
		},
	)
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

func cacheHandlerWithBody[In any, Out any](
	c *gin.Context,
	ttl time.Duration,
	db *redis.Client,
	fetch func(*In) (*Out, error),
) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	key := getSHA256Hash(c.Request.URL.Path + string(body))
	dataInCache := cache.CheckCacheHit[Out](key, db)
	if dataInCache != nil {
		c.Render(http.StatusOK, render.JSON{Data: dataInCache})
		return
	}

	var input In
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	resp, err := fetch(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if jsonData, err := goccyjson.Marshal(resp); err == nil {
		db.Set(key, jsonData, ttl)
	}

	c.Render(http.StatusOK, render.JSON{Data: resp})
}
