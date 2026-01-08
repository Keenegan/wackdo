package controllers_user

import (
	"net/http"
	"net/mail"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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

	if _, err := mail.ParseAddress(body.Email); err != nil {
		c.Error(&service.InvalidParamError{
			Reason: "email is invalid",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.Error(err)
		return
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
		Role:     models.RoleEmployee,
	}

	_, err = service.CreateUser(user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
