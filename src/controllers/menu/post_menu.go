package controllers_menus

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

type MenuPostRequest struct {
	Name        string  `json:"name" binding:"required"`
	BasePrice   float32 `json:"basePrice" binding:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	ProductIds  []uint  `json:"productIds"`
}

func PostMenu(c *gin.Context) {
	var req MenuPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ValidateMenuPostRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if exists, err := service.MenuExists(req.Name); exists || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu name already exists"})
		return
	}

	products, err := service.GetProductsByIds(req.ProductIds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(products) != len(req.ProductIds) {
		// todo say which
		c.JSON(http.StatusBadRequest, gin.H{"error": "Some products ids were not found"})
		return
	}
	for _, product := range products {
		if !product.Available {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("product %d is not available", product.ID),
			})
			return
		}
	}

	newMenu, err := service.CreateMenu(models.Menu{
		Name:        req.Name,
		BasePrice:   req.BasePrice,
		Description: req.Description,
		Image:       req.Image,
		Products:    products,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          newMenu.ID,
		"name":        newMenu.Name,
		"basePrice":   newMenu.BasePrice,
		"description": newMenu.Description,
		"image":       newMenu.Image,
		"productIds":  newMenu.Products,
	})
}

func ValidateMenuPostRequest(req *MenuPostRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if req.BasePrice <= 0 {
		return errors.New("basePrice must be greater than 0")
	}

	if len(req.ProductIds) == 0 {
		return errors.New("productIds must contain at least one product")

	}

	return nil
}
