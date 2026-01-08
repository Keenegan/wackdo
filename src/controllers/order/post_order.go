package controllers_order

import (
	"net/http"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

type OrderPostRequest struct {
	Items []service.OrderItemRequest `json:"items" binding:"required"`
}

func PostOrder(c *gin.Context) {
	var req OrderPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "order must contain at least one item",
		})
		return
	}

	order, err := service.CreateOrder(req.Items)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         order.ID,
		"status":     order.Status,
		"totalPrice": order.TotalPrice,
		"orderLines": order.OrderLines,
		"createdAt":  order.CreatedAt,
	})
}
