package models

import (
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
	var dsn = "user:test@tcp(product-db:3306)" + "/" + "product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}
