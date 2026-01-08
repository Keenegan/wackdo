package controllers_menus

import (
	"net/http"
	"strconv"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

// todo get only available products
func GetMenu(c *gin.Context) {
	var menus = []models.Menu{}

	name := c.Query("name")
	idStr := c.Query("id")

	// Find menu by ID
	if idStr != "" && name != "" {
		c.Error(&service.InvalidParamError{
			Reason: "cannot filter by id and name at the same time",
		})
		return

	} else if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			c.Error(&service.InvalidParamError{
				Reason: "id must be a number greater than 0",
			})
			return
		}

		menu, err := service.GetMenuById(id)
		if err != nil {
			c.Error(err)
			return
		}
		menus = append(menus, menu)

	} else if name != "" {
		// Find menu by name
		menu, err := service.GetMenuByName(name)
		if err != nil {
			c.Error(err)
			return
		}

		menus = append(menus, menu)

	} else {
		// Default: find all menus paginated
		pageSize := 10
		pageNum := service.GetPagerFromContext(c)
		if pageNum < 1 {
			pageNum = 1
		}
		page := (pageNum - 1) * pageSize

		results, err := service.GetMenus(page, pageSize)
		if err != nil {
			c.Error(err)
			return
		}
		menus = append(menus, results...)

	}

	c.JSON(http.StatusOK, menus)
}

// todo add test for this file

