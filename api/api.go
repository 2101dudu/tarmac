package api

import (
	"net/http"
	"strconv"
	"tarmac/internal/render"
	"tarmac/wsdl"
	"time"

	"github.com/gin-gonic/gin"
)

type Api struct {
	SoapService wsdl.Wbs_pkt_methodsSoap
	Credentials *wsdl.CredentialsStruct
}

func (a *Api) Start() {
	engine := gin.Default()

	engine.GET("api/search/product", a.handleSearchProduct)
	engine.GET("api/get/product/:prodCode", a.handleDynGetProductParameters)

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	srv.ListenAndServe()
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
