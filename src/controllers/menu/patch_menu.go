package controllers_menus

import (
	"fmt"
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
	ProductIds  []uint  `json:"productIds"`
}

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

	if len(req.ProductIds) > 0 {
		products, err := service.GetProductsByIds(req.ProductIds)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(products) != len(req.ProductIds) {
			foundIds := make(map[uint]bool)
			for _, p := range products {
				foundIds[p.ID] = true
			}
			var missingIds []uint
			for _, id := range req.ProductIds {
				if !foundIds[id] {
					missingIds = append(missingIds, id)
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error":      "Some product IDs were not found",
				"missingIds": missingIds,
			})
			return
		}
		for _, product := range products {
			if !product.Available {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("product %d is not available", product.ID),
				})
				return
			}
		}
		menu.Products = products
	}

	menu, err = service.UpdateMenu(menu)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, menu)
}
