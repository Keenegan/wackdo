package controllers_products

import (
	"net/http"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

type ProductDeleteRequest struct {
	ID int `json:"id" binding:"required"`
}

// todo delete product in a menu ?
func DeleteProduct(c *gin.Context) {
	var req ProductDeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := service.DeleteProductById(req.ID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, "")
}

// todo add test for this file
