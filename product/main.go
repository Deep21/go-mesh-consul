package main

import (
	"os"

	"deep21/go-mesh-consul/product/models"
	"deep21/go-mesh-consul/product/route"
)

func initialMigration() {
	models.DB.AutoMigrate(&models.Product{})
}

func main() {
	models.ConnectDatabase()

	initialMigration()
	r := route.SetupRouter()
	var appPort string = os.Getenv("APP_PORT")

	r.Run(":" + appPort)
}
