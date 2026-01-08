package controllers_user

import (
	"net/http"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	page := service.GetPagerFromContext(c)
	pageSize := 20

	users, err := service.GetUsers(page-1, pageSize)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"page":  page,
	})
}
