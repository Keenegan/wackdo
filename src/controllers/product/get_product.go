package controllers_products

import (
	"net/http"
	"wackdo/src/service"
	product_repository "wackdo/src/service/repository"

	"github.com/gin-gonic/gin"
)

type ProductGetRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetProducts(c *gin.Context) {
	var req ProductGetRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// Find all products paginated
		pageSize := 10
		page := (service.GetPagerFromContext(c) - 1) * pageSize

		if products, err := product_repository.GetProducts(page, pageSize); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusCreated, products)

		}
		return
	}
	if req.ID != 0 {
		// Find product by ID
		if product, err := product_repository.GetProductById(int(req.ID)); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusCreated, product)
		}
	}
	if req.Name != "" {
		// Find product by name
		if product, err := product_repository.GetProductByName(req.Name); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusCreated, product)
		}
	}

}
