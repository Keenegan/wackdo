package controllers_order

import (
	"net/http"
	"strconv"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

type OrderPatchRequest struct {
	Status models.OrderStatus `json:"status" binding:"required"`
}

func PatchOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order ID",
		})
		return
	}

	var req OrderPatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !req.Status.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order status",
		})
		return
	}

	order, err := service.UpdateOrderStatus(uint(id), req.Status)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        order.ID,
		"status":    order.Status,
		"updatedAt": order.UpdatedAt,
	})
}
