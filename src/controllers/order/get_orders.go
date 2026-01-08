package controllers_order

import (
	"net/http"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	page := service.GetPagerFromContext(c)
	pageSize := 20

	orders, err := service.GetOrders(page-1, pageSize)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"page":   page,
	})
}
