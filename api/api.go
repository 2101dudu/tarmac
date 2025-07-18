package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"tarmac/coordinator"
	"tarmac/utils"
	"tarmac/wsdl"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type Api struct {
	Coordinator         *coordinator.Service
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
	engine.POST("api/search/product", a.handleSearchProductWithBody)
	engine.GET("api/search/product/page", a.handleSearchProductPagination)
	engine.GET("api/get/product/:prodCode", a.handleDynGetProductParameters)
	engine.POST("api/dynamic/product/available-services", a.handleDynSearchProductAvailableServices)
	engine.GET("api/dynamic/product/available-services/status", a.handleAsyncAvailableServicesStatus)
	engine.GET("api/dynamic/product/available-services/flights/page", a.handleDynSearchProductAvailableServicesFlightsPagination)
	engine.GET("api/dynamic/product/available-services/hotels/page", a.handleDynSearchProductAvailableServicesHotelsPagination)
	engine.GET("api/dynamic/product/available-services/hotels/:hotelCode/rooms/page", a.handleDynSearchProductAvailableServicesHotelRoomsPagination)
	engine.POST("api/dynamic/product/set-services", a.handleDynSetServicesSelectedAndGetToken)
	engine.GET("api/dynamic/product/get-simulation", a.handleDynGetSimulation)
	engine.POST("api/dynamic/product/send-email", a.handleSendEmail)
	engine.GET("api/page/highlighted/tag", a.handleGetHighlightedTag)

	//========================================================================
	engine.POST("api/admin/auth", a.handleAdminAuth) // ADMIN ENDPOINT

	admin := engine.Group("/api/admin")
	admin.Use(a.AdminAuthMiddleware())
	admin.GET("/products", a.handleListAllProducts)
	admin.POST("/products/:prodCode/tags/add", a.handleAddProductTags)
	admin.GET("/products/:prodCode/tags/remove/:tagToRemove", a.handleRemoveProductTags)
	admin.POST("/products/:prodCode/toggle", a.handleProductToggleState)
	admin.GET("/page/highlight/:tagToHighlight", a.handleHighlightTag)

	//========================================================================

	//========================================================================
	engine.GET("api/admin/reset-state", a.handleResetState) // DEV PURPOSE
	//========================================================================

	//========================================================================
	engine.GET("api/admin/sync/products", a.handleSyncAllProducts) // MANUAL CALL

	c := cron.New()
	c.AddFunc("0 */6 * * *", func() {
		_ = a.Coordinator.HandleSyncAllProducts()
	})
	c.Start()
	//========================================================================

	engine.Run("192.168.1.120:8080")
}

func (a *Api) handleGetMasterData(c *gin.Context) {
	data, err := a.Coordinator.HandleGetMasterData(c.Request.URL.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (a *Api) handleSearchProductWithBody(c *gin.Context) {
	var query coordinator.ProductQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	firstPage, token, hasMore, err := a.Coordinator.HandleSearchProductWithBody(query, c.Request.URL.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": firstPage, "token": token, "hasMore": hasMore})
}

func (a *Api) handleSearchProductPagination(c *gin.Context) {
	token, cursor, limit, err := parsePaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, hasMore, err := a.Coordinator.HandleSearchProductPagination(token, cursor, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": page, "hasMore": hasMore})
}

func (a *Api) handleDynGetProductParameters(c *gin.Context) {
	prodCodeRaw := c.Param("prodCode")
	data, err := a.Coordinator.HandleDynGetProductParameters(prodCodeRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (a *Api) handleDynSearchProductAvailableServices(c *gin.Context) {
	var input wsdl.DynProductAvailableServicesRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	searchID := a.Coordinator.HandleDynSearchProductAvailableServices(input)
	c.JSON(http.StatusOK, gin.H{"searchId": searchID})
}

func (a *Api) handleAsyncAvailableServicesStatus(c *gin.Context) {
	searchID := c.Query("id")
	if searchID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Missing search id")})
		return
	}

	status, resp, flightsToken, hasMoreFlights, hotelsToken, hasMoreHotels, err := a.Coordinator.HandleAsyncAvailableServicesStatus(searchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "data": resp, "flightsToken": flightsToken, "hasMoreFlights": hasMoreFlights, "hotelsToken": hotelsToken, "hasMoreHotels": hasMoreHotels})
}

func (a *Api) handleDynSearchProductAvailableServicesFlightsPagination(c *gin.Context) {
	token, cursor, limit, err := parsePaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flights, hasMore, err := a.Coordinator.HandleDynSearchProductAvailableServicesFlightsPagination(token, cursor, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"flights": flights, "hasMore": hasMore})
}

func (a *Api) handleDynSearchProductAvailableServicesHotelsPagination(c *gin.Context) {
	token, cursor, limit, err := parsePaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotels, hasMore, err := a.Coordinator.HandleDynSearchProductAvailableServicesHotelsPagination(token, cursor, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"hotels": hotels, "hasMore": hasMore})
}

func (a *Api) handleDynSearchProductAvailableServicesHotelRoomsPagination(c *gin.Context) {}

func (a *Api) handleDynSetServicesSelectedAndGetToken(c *gin.Context) {
	prodCode := c.Query("prodCode")
	var input wsdl.DynServicesSelectedRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.Coordinator.HandleDynSetServicesSelectedAndGetToken(input, prodCode)
	if err != nil || token == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": *token})
}

func (a *Api) handleDynGetSimulation(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Missing token")})
		return
	}
	simul, err := a.Coordinator.HandleDynGetSimulation(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, simul)
}

func (a *Api) handleSendEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Missing token")})
		return
	}
	var contact coordinator.ContactInfo
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	simul, err := a.Coordinator.HandleSendEmail(token, contact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, simul)
}

/*
func (a *Api) handleDynCheckSetServices(c *gin.Context) {
	var input wsdl.DynGetOptionalsSelectedRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	resp, err := a.WSDLService.DynGetOptionalsSelected(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Render(http.StatusOK, render.JSON{Data: resp})
}
*/
/*
func (a *Api) handleDynGetSimulation(c *gin.Context) {
	var input wsdl.DynGetSimulationRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	resp, err := a.WSDLService.DynGetSimulation(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Render(http.StatusOK, render.JSON{Data: resp})
}
*/

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
	go a.Coordinator.HandleResetState()
	c.JSON(http.StatusOK, gin.H{"status": "reset concluded"})
}

func (a *Api) handleSyncAllProducts(c *gin.Context) {
	err := a.Coordinator.HandleSyncAllProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product sync ended."})
}

func (a *Api) handleAdminAuth(c *gin.Context) {
	var body struct{ Password string }
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expected := a.AdminHashedPassword
	if utils.GetSHA256Hash(body.Password) != expected {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token := a.Coordinator.HandleAdminAuth()
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
		if storedToken := a.Coordinator.HandleCheckAdminAuth(token); storedToken == nil || *storedToken != "valid" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func (a *Api) handleListAllProducts(c *gin.Context) {
	productList, err := a.Coordinator.HandleGetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB failed: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"list": productList})
}

func (a *Api) handleAddProductTags(c *gin.Context) {
	prodCodeRaw := c.Param("prodCode")
	_, err := strconv.Atoi(prodCodeRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req struct {
		Tags []string `json:"tags"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go a.Coordinator.HandleAddProductTags(prodCodeRaw, req.Tags)
	c.JSON(http.StatusOK, gin.H{"message": "Tags added successfully"})
}

func (a *Api) handleRemoveProductTags(c *gin.Context) {
	tagToRemove, prodCodeRaw := c.Param("tagToRemove"), c.Param("prodCode")
	_, err := strconv.Atoi(prodCodeRaw)
	if err != nil || tagToRemove == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Incorrect parameters")})
		return
	}

	go a.Coordinator.HandleRemoveProductTags(tagToRemove, prodCodeRaw)
	c.JSON(http.StatusOK, gin.H{"message": "Tag removed successfully"})
}

func (a *Api) handleProductToggleState(c *gin.Context) {
	prodCodeRaw := c.Param("prodCode")
	_, err := strconv.Atoi(prodCodeRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go a.Coordinator.HandleProductToggleState(prodCodeRaw)
	c.JSON(http.StatusOK, gin.H{"message": "Product state changed successfully"})
}

func (a *Api) handleHighlightTag(c *gin.Context) {
	tagName := c.Param("tagToHighlight")
	go a.Coordinator.HandleHighlightTag(tagName)
	c.JSON(http.StatusOK, gin.H{"message": "Product enabled successfully"})
}

func (a *Api) handleGetHighlightedTag(c *gin.Context) {
	tagName, _ := a.Coordinator.HandleGetHighlightedTag()
	if tagName == nil {
		c.JSON(http.StatusOK, gin.H{"tag": "charter"}) // default
		return
	}
	c.JSON(http.StatusOK, gin.H{"tag": *tagName})
}
