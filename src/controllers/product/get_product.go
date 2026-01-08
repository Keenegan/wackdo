package controllers_products

import (
	"net/http"
	"strconv"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

// todo get only available products
func GetProducts(c *gin.Context) {
	var products = []models.Product{}

	name := c.Query("name")
	idStr := c.Query("id")

	// Find product by ID
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

		product, err := service.GetProductById(id)
		if err != nil {
			c.Error(err)
			return
		}
		products = append(products, product)

	} else if name != "" {
		// Find product by name
		product, err := service.GetProductByName(name)
		if err != nil {
			c.Error(err)
			return
		}

		products = append(products, product)

	} else {
		// Default: find all products paginated
		pageSize := 10
		pageNum := service.GetPagerFromContext(c)
		if pageNum < 1 {
			pageNum = 1
		}
		page := (pageNum - 1) * pageSize

		results, err := service.GetProducts(page, pageSize)
		if err != nil {
			c.Error(err)
			return
		}
		products = append(products, results...)

	}

	c.JSON(http.StatusOK, products)
}
// todo add test for this file
