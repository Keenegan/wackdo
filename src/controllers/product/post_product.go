package controllers_products

import (
	"errors"
	"net/http"
	"strings"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

type ProductPostRequest struct {
	Name        string          `json:"name" binding:"required"`
	BasePrice   float32         `json:"basePrice" binding:"required"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	Category    models.Category `json:"category" binding:"required"`
	Available   *bool           `json:"available"`
}

func PostProduct(c *gin.Context) {
	var req ProductPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ValidateProductPostRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if exists, err := service.ProductExists(req.Name); exists || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name already exists"})
		return
	}
	newProduct, err := service.CreateProduct(models.Product{
		Name:        req.Name,
		BasePrice:   req.BasePrice,
		Description: req.Description,
		Image:       req.Image,
		Category:    req.Category,
		Available:   *req.Available,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          newProduct.ID,
		"name":        newProduct.Name,
		"basePrice":   newProduct.BasePrice,
		"description": newProduct.Description,
		"image":       newProduct.Image,
		"category":    newProduct.Category,
		"available":   newProduct.Available,
	})
}

func ValidateProductPostRequest(req *ProductPostRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if req.Available == nil {
		defaultAvailable := true
		req.Available = &defaultAvailable
	}

	if req.BasePrice <= 0 {
		return errors.New("basePrice must be greater than 0")
	}

	if !req.Category.IsValid() {
		return errors.New("invalid category")
	}

	return nil
}

