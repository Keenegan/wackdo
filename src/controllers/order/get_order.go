package controllers_order

import (
	"net/http"
	"strconv"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order ID",
		})
		return
	}

	order, err := service.GetOrderById(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         order.ID,
		"status":     order.Status,
		"totalPrice": order.TotalPrice,
		"orderLines": order.OrderLines,
		"createdAt":  order.CreatedAt,
		"updatedAt":  order.UpdatedAt,
	})
}
