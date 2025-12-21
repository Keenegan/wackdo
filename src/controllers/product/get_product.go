package controllers_products

import (
	"net/http"
	"strconv"
	"wackdo/src/service"
	product_repository "wackdo/src/service/repository"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	// Find product by ID
	idStr := c.Query("id")
	if idStr != "" && idStr != "0" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		product, err := product_repository.GetProductById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, product)
		return
	}

	// Find product by name
	name := c.Query("name")
	if name != "" {
		product, err := product_repository.GetProductByName(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, product)
		return
	}

	// Default: find all products paginated
	pageSize := 10
	pageNum := service.GetPagerFromContext(c)
	if pageNum < 1 {
		pageNum = 1
	}
	page := (pageNum - 1) * pageSize

	products, err := product_repository.GetProducts(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
