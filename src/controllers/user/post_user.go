package controllers_user

import (
	"net/http"
	"wackdo/src/initializers"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// todo add error message if email exists
func Register(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.BindJSON(&body); err != nil {
		c.Error(&service.InvalidParamError{
			Reason: err.Error(),
		})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
		Role:     models.RoleEmployee,
	}

	initializers.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
