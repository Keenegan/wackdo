package middleware

import (
	"errors"
	"net/http"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var invalidParamErr *service.InvalidParamError
		var duplicateEmailErr *service.DuplicateEmailError
		var entityNotFoundErr *service.EntityNotFoundError

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		status := http.StatusInternalServerError
		message := "internal server error"

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			status = http.StatusNotFound
			message = "resource not found"

		case errors.As(err, &entityNotFoundErr):
			status = http.StatusNotFound
			message = entityNotFoundErr.Error()

		case errors.As(err, &invalidParamErr):
			status = http.StatusBadRequest
			message = invalidParamErr.Error()

		case errors.As(err, &duplicateEmailErr):
			status = http.StatusConflict
			message = duplicateEmailErr.Error()
		}
		c.JSON(status, gin.H{
			"error": message,
		})
	}
}
