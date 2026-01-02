package controllers_menus

import (
	"net/http"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"github.com/gin-gonic/gin"
)

type MenuUpdateRequest struct {
	ID          uint    `json:"id" binding:"required"`
	Name        string  `json:"name"`
	BasePrice   float32 `json:"basePrice"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

// todo allow product id change
func UpdateMenu(c *gin.Context) {
	var req MenuUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var menu models.Menu
	if err := initializers.DB.First(&menu, req.ID).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if req.Name != menu.Name {
		var count int64
		initializers.DB.Model(&models.Menu{}).
			Where("name = ?", req.Name).
			Count(&count)

		if count > 0 {
			c.JSON(400, gin.H{"error": "Menu name already exists"})
			return
		}
	}

	menu.Name = req.Name
	menu.BasePrice = req.BasePrice
	menu.Description = req.Description
	menu.Image = req.Image

	err := initializers.DB.Save(&menu).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}
