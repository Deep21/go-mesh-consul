package controller

import (
	"net/http"

	"github.com/deep21/go-mesh-consul/product/models"

	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {
	var p models.Product

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product = models.Product{
		Reference:   p.Reference,
		ProductName: p.ProductName,
		Description: p.Description,
	}

	models.DB.Create(&product)

}

func NewProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func AllProducts() *[]models.Product {
	var users []models.Product
	models.DB.Find(&users)
	return &users
}

func FindProduct(id string) *models.Product {
	var product models.Product

	models.DB.Find(&product)
	return &product
}

func DeleteUser(c *gin.Context) {
	var u models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&u)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
