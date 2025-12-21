package controllers_products

import (
	"net/http"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"github.com/gin-gonic/gin"
)

type ProductDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

func DeleteProduct(c *gin.Context) {
	var req ProductDeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := initializers.DB.Delete(&models.Product{}, req.ID).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, "")
}
