package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Product struct {
	gorm.Model
	ID          int
	Reference   string `json:"reference" binding:"required"`
	ProductName string `json:"productname" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint8  `json:"price" binding:"required"`
}

func ConnectDatabase() {
	var port string = os.Getenv("PORT")
	var user string = os.Getenv("USER")
	var pwd string = os.Getenv("PWD")
	var ip string = "172.21.0.4:" + port
	var dsn = user + ":" + pwd + "@tcp(" + ip + ")" + "/" + "product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}
