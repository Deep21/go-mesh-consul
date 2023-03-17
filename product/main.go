package main

import (
	"deep21/go-mesh-consul/product/models"
	"deep21/go-mesh-consul/product/route"
	"fmt"
	"os"
)

func initialMigration() {
	models.DB.AutoMigrate(&models.Product{})
}

func main() {
	models.ConnectDatabase()

	// initialMigration()
	var p []models.Product
	models.DB.Find(&p)
	fmt.Println(p)

	r := route.SetupRouter()
	var appPort string = os.Getenv("APP_PORT")

	r.Run(":" + appPort)
}
