package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tarmac/api/helper"
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
	key := getSHA256Hash(c.Request.URL.String())

	value, err := a.DBService.Get(key).Result()
	if err == nil { // has entry, return cache
		var cachedResp *wsdl.SearchProductMasterDataResponse
		if err := goccyjson.Unmarshal([]byte(value), &cachedResp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Corrupted cache data"})
			return
		}
		dbSize, err := a.DBService.DBSize().Result()
		if err != nil {
			log.Fatalf("failed to get DB size: %v", err)
		}
		fmt.Printf("Redis DB has %d keys\n", dbSize)

		c.Render(http.StatusOK, render.JSON{Data: cachedResp})
		return
	}

	masterData, err := a.SoapService.SearchProductMasterData(&wsdl.SearchProductMasterDataRequest{
		Credentials: a.Credentials,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get master data: " + err.Error()})
		return
	}

	jsonData, err := goccyjson.Marshal(masterData)
	if err != nil {
		log.Fatalf("failed to marshal trips: %v", err)
	}

	a.DBService.Set(key, jsonData, 1*time.Hour)

	c.Render(http.StatusOK, render.JSON{Data: masterData})
}

func (a *Api) handleSearchProduct(c *gin.Context) {
	productSearch, err := a.SoapService.SearchProducts(&wsdl.SearchProductRequest{
		Credentials: a.Credentials,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products: " + err.Error()})
		return
	}

	c.Render(http.StatusOK, render.JSON{Data: productSearch})
}

func (a *Api) handleSearchProductWithBody(c *gin.Context) {
	var input wsdl.SearchProductRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	input.Credentials = a.Credentials
	response, err := a.SoapService.SearchProducts(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products: " + err.Error()})
		return
	}

	err = helper.OptimizeProductSearch(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid number of products: " + err.Error()})
		return
	}

	c.Render(http.StatusOK, render.JSON{Data: response})
}

func (a *Api) handleDynGetProductParameters(c *gin.Context) {
	prodCode := c.Param("prodCode")
	if _, err := strconv.Atoi(prodCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code must be a number"})
		return
	}

	productParams, err := a.SoapService.DynGetProductParameters(&wsdl.DynProductParametersRequest{
		Credentials: a.Credentials,
		Productcode: &prodCode,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product parameters: " + err.Error()})
		return
	}

	c.Render(http.StatusOK, render.JSON{Data: productParams})
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
