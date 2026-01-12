package controllers_products

import (
	"net/http"
	"strconv"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id parameter",
		})
		return
	}

	if err := service.DeleteProductById(id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, "")
}
// todo add test for this file
