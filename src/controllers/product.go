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

	var count int64
	initializers.DB.Model(&models.Product{}).
		Where("name = ?", req.Name).
		Count(&count)

	if count > 0 {
		c.JSON(400, gin.H{"error": "Product name already exists"})
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
		c.JSON(400, gin.H{"error": err.Error()})
	}

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

type ProductGetRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetProducts(c *gin.Context) {
	var req ProductGetRequest
	var products []models.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		// Find all products paginated
		pageSize := 10
		page := (GetPagerFromContext(c) - 1) * pageSize

		if err := initializers.DB.Order("id ASC").Limit(pageSize).Offset(page).Find(&products).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	} else if req.ID != 0 {
		// Find product by ID
		if err := initializers.DB.Where("id = ?", req.ID).Find(&products).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	} else if req.Name != "" {
		// Find product by name
		if err := initializers.DB.Where("name = ?", req.Name).Find(&products).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, products)

}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
