package api

import (
	"net/http"
	"strconv"
	"tarmac/internal/render"
	"tarmac/wsdl"

	"github.com/gin-gonic/gin"
)

type Api struct {
	SoapService wsdl.Wbs_pkt_methodsSoap
	Credentials *wsdl.CredentialsStruct
}

func (a *Api) Start() {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)

	engine.GET("api/search/product", a.handleSearchProduct)
	engine.GET("api/get/product/:prodCode", a.handleDynGetProductParameters)
	engine.POST("api/dynamic/product/available-services", a.handleDynSearchProductAvailableServices)
	engine.POST("api/dynamic/product/set-services", a.handleDynSetServicesSelected)
	engine.POST("api/dynamic/product/check-set-services", a.handleDynCheckSetServices)
	engine.POST("api/dynamic/product/get-simulation", a.handleDynGetSimulation)

	engine.Run(":8080")
}

func (a *Api) handleSearchProduct(c *gin.Context) {
	productSearch, err := a.SoapService.SearchProducts(&wsdl.SearchProductRequest{
		Credentials: a.Credentials,
	})
	HandleErr(err)

	c.Render(http.StatusOK, render.JSON{Data: productSearch})
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
	HandleErr(err)

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
