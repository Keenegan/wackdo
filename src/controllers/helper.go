package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPagerFromContext(c *gin.Context) (page int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return 1
	}
	if page < 1 {
		page = 1
	}
	return
}
