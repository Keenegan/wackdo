package controllers_products

import (
	"net/http"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

type ProductUpdateRequest struct {
	ID          uint            `json:"id" binding:"required"`
	Name        string          `json:"name"`
	BasePrice   float32         `json:"basePrice"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	Category    models.Category `json:"category"`
	Available   *bool           `json:"available"`
}

func UpdateProduct(c *gin.Context) {
	var req ProductUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := service.GetProductById(int(req.ID))
	if err != nil {
		c.Error(err)
		return
	}

	if req.Name != product.Name {
		exists, err := service.ProductExists(req.Name)
		if err != nil {
			c.Error(err)
			return
		}

		if exists {
			c.Error(&service.InvalidParamError{
				Reason: "Product name already exists",
			})
			return
		}
	}

	product.Name = req.Name
	product.BasePrice = req.BasePrice
	product.Description = req.Description
	product.Category = req.Category
	product.Image = req.Image

	if req.Available != nil {
		product.Available = *req.Available
	}

	product, err = service.UpdateProduct(product)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, product)
}
