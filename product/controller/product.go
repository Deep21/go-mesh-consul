package controller

import (
	"errors"
	"net/http"

	"deep21/go-mesh-consul/product/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProduct(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"data": true})

}

func AllProducts() *[]models.Product {
	var p []models.Product
	models.DB.Find(&p)
	return &p
}

func FindProduct(c *gin.Context) *models.Product {
	var product models.Product
	var id string = c.Param("id")

	r := models.DB.First(&product, "id = ?", id)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return nil
	}
	return &product
}

func DeleteUser(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
