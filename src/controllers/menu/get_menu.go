package controllers_menus

import (
	"net/http"
	"strconv"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

// GetMenus returns all menus with pagination
func GetMenus(c *gin.Context) {
	pageSize := 10
	pageNum := service.GetPagerFromContext(c)
	if pageNum < 1 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize

	menus, err := service.GetMenus(offset, pageSize)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, menus)
}

// GetMenuById returns a menu by its ID
func GetMenuById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Error(&service.InvalidParamError{
			Reason: "id must be a positive number",
		})
		return
	}

	menu, err := service.GetMenuById(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, menu)
}

// GetMenuByName searches for a menu by name
func GetMenuByName(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		c.Error(&service.InvalidParamError{
			Reason: "name query parameter is required",
		})
		return
	}

	menu, err := service.GetMenuByName(name)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, menu)
}
