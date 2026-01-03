package controllers_menus

import (
	"net/http"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

type MenuDeleteRequest struct {
	ID int `json:"id" binding:"required"`
}

func DeleteMenu(c *gin.Context) {
	var req MenuDeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := service.DeleteMenuById(req.ID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, "")
}

// todo add test for this file
