package controllers_products

import (
	"errors"
	"net/http"
	"strings"
	"wackdo/src/initializers"
	"wackdo/src/models"

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

	var count int64
	initializers.DB.Model(&models.Product{}).
		Where("name = ?", req.Name).
		Count(&count)

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name already exists"})
		return
	}

	insertProduct := &models.Product{
		Name:        req.Name,
		BasePrice:   req.BasePrice,
		Description: req.Description,
		Image:       req.Image,
		Category:    req.Category,
		Available:   *req.Available,
	}

	err := initializers.DB.Create(insertProduct).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          insertProduct.ID,
		"name":        insertProduct.Name,
		"basePrice":   insertProduct.BasePrice,
		"description": insertProduct.Description,
		"image":       insertProduct.Image,
		"category":    insertProduct.Category,
		"available":   insertProduct.Available,
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

// todo use repository