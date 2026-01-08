package controllers_menus

import (
	"net/http"
	"wackdo/src/service"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	menu, err := service.GetMenuById(int(req.ID))
	if err != nil {
		c.Error(err)
		return
	}

	if req.Name != menu.Name {
		exists, err := service.MenuExists(req.Name)
		if err != nil {
			c.Error(err)
			return
		}

		if exists {
			c.Error(&service.InvalidParamError{
				Reason: "Menu name already exists",
			})
			return
		}
	}

	menu.Name = req.Name
	menu.BasePrice = req.BasePrice
	menu.Description = req.Description
	menu.Image = req.Image

	menu, err = service.UpdateMenu(menu)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, menu)
}
