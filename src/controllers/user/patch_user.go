package controllers_user

import (
	"net/http"
	"net/mail"
	"strconv"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// todo check
type UserUpdateRequest struct {
	Email    *string      `json:"email"`
	Password *string      `json:"password"`
	Role     *models.Role `json:"role"`
}

func UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var body UserUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if _, err := mail.ParseAddress(*body.Email); err != nil {
		c.Error(&service.InvalidParamError{
			Reason: "email is invalid",
		})
		return
	}

	user, err := service.GetUserById(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if body.Email != nil {
		user.Email = *body.Email
	}

	if body.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*body.Password), 10)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "password error"})
			return
		}
		user.Password = string(hash)
	}

	if body.Role != nil {
		if !models.IsValidRole(*body.Role) && *body.Role != models.RoleAdmin {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
			return
		}

		user.Role = *body.Role
	}

	user, err = service.UpdateUserFull(user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
		"user":    user,
	})
}
