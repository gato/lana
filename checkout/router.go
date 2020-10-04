package checkout

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddRoutes - add routes for basket and checkout management
func AddRoutes(rg *gin.RouterGroup) {

	r := rg.Group("/basket")

	r.GET("/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		HandleGetByID(c, id)
	})

	r.POST("/", func(c *gin.Context) {
		HandleCreateEmtpyBasket(c)
	})

	r.DELETE("/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		HandleDeleteBasket(c, id)
	})

	// Route add product
	r.POST("/:id", func(c *gin.Context) {
		var _item ProductItem
		id := c.Params.ByName("id")
		if err := c.BindJSON(&_item); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		HandleAddProduct(c, id, _item)
	})
}
