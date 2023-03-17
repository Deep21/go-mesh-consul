package route

import (
	"net/http"

	"deep21/go-mesh-consul/product/controller"
	"deep21/go-mesh-consul/product/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/products", func(ctx *gin.Context) {
			var sp []models.Product
			models.DB.Find(&sp)
			ctx.IndentedJSON(http.StatusOK, sp)
		})

		v1.GET("/product/:id", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, controller.FindProduct("3"))
		})

		v1.POST("/products", controller.NewProduct)

		v1.DELETE("/product/:id", controller.DeleteUser)
	}
	return r
}
