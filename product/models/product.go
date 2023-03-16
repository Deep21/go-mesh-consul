package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Product struct {
	gorm.Model
	ID          int
	Reference   string
	ProductName string
	Description string
	Price       uint8
}

func ConnectDatabase() {
	// var ip string = "product-db:3306"
	var dsn = "user:test@tcp(product-db:3306)" + "/" + "product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}
