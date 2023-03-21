package route

import (
	"io"
	"net/http"
	"os"

	"deep21/go-mesh-consul/product/controller"
	"deep21/go-mesh-consul/product/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	v1 := r.Group("/v1")
	{
		v1.GET("/products", func(ctx *gin.Context) {
			var sp *[]models.Product = controller.AllProducts()
			models.DB.Find(&sp)
			ctx.IndentedJSON(http.StatusOK, sp)
		})

		v1.GET("/product/:id", func(ctx *gin.Context) {
			var sp *models.Product = controller.FindProduct(ctx)
			ctx.IndentedJSON(http.StatusOK, *sp)
		})

		v1.POST("/products", controller.NewProduct)

		v1.DELETE("/product/:id", controller.DeleteUser)
	}
	return r
}
