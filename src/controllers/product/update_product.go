package controllers_products

import (
	"net/http"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"github.com/gin-gonic/gin"
)

type ProductUpdateRequest struct {
	ID          uint            `json:"id" binding:"required"`
	Name        string          `json:"name"`
	BasePrice   float32         `json:"basePrice"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	Category    models.Category `json:"category"`
	Available   *bool           `json:"available"`
}

func UpdateProduct(c *gin.Context) {
	var req ProductUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var product models.Product
	if err := initializers.DB.First(&product, req.ID).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if req.Name != product.Name {
		var count int64
		initializers.DB.Model(&models.Product{}).
			Where("name = ?", req.Name).
			Count(&count)

		if count > 0 {
			c.JSON(400, gin.H{"error": "Product name already exists"})
			return
		}
	}

	product.Name = req.Name
	product.BasePrice = req.BasePrice
	product.Description = req.Description
	product.Category = req.Category
	product.Image = req.Image

	if req.Available != nil {
		product.Available = *req.Available
	}

	err := initializers.DB.Save(&product).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}
