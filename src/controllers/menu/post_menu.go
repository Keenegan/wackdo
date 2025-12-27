package controllers_menus

import (
	"github.com/gin-gonic/gin"
)

// todo add menu options
type MenuPostRequest struct {
	Name        string          `json:"name" binding:"required"`
	BasePrice   float32         `json:"basePrice" binding:"required"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
}

func PostMenu(c *gin.Context) {

}
