package main

import (
	"deep21/go-mesh-consul/product/models"
	"deep21/go-mesh-consul/product/route"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var appPort string = os.Getenv("APP_PORT")

func initialMigration() {
	models.DB.AutoMigrate(&models.Product{})
}

func main() {
	models.ConnectDatabase()

	initialMigration()
	var p []models.Product
	models.DB.Find(&p)
	fmt.Println(p)

	r := route.SetupRouter()
	r.Use(CORSMiddleware())
	r.Run(":" + appPort)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
