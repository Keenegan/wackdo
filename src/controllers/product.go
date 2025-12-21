package controllers

import (
	"errors"
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
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ValidateProductPostRequest(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
	initializers.DB.Create(insertProduct)

	c.JSON(200, gin.H{
		"name":        req.Name,
		"basePrice":   req.BasePrice,
		"description": req.Description,
		"image":       req.Image,
		"category":    req.Category,
		"available":   req.Available,
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
		return errors.New("base price can't be 0 or less")
	}

	if !req.Category.IsValid() {
		return errors.New("invalid category")
	}

	return nil
}
