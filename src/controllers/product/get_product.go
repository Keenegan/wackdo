package controllers_products

import (
	"net/http"
	"strconv"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

// GetProducts returns all products with pagination
func GetProducts(c *gin.Context) {
	pageSize := 10
	pageNum := service.GetPagerFromContext(c)
	if pageNum < 1 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize

	products, err := service.GetProducts(offset, pageSize)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductById returns a product by its ID
func GetProductById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.Error(&service.InvalidParamError{
			Reason: "id must be a positive number",
		})
		return
	}

	product, err := service.GetProductById(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProductByName searches for products by name (partial match)
func GetProductByName(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		c.Error(&service.InvalidParamError{
			Reason: "name query parameter is required",
		})
		return
	}

	products, err := service.GetProductByName(name)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, products)
}
