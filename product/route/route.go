package route

import (
	"net/http"

	"github.com/deep21/go-mesh-consul/product/controller"

	"github.com/deep21/go-mesh-consul/product/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/products", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, controller.AllProducts())
		})

		v1.GET("/product/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			var f *models.Product = controller.FindProduct(id)
			ctx.IndentedJSON(http.StatusOK, f)
		})

		v1.POST("/products", controller.NewProduct)

		v1.DELETE("/product/:id", controller.DeleteUser)
	}
	return r
}
