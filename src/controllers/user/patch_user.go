package controllers_user

import (
	"net/http"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserUpdateRequest struct {
	Email    *string      `json:"email"`
	Password *string      `json:"password"`
	Role     *models.Role `json:"role"`
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var body UserUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	requesterRole := c.MustGet("role").(models.Role)

	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
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
		if !models.IsValidRole(*body.Role) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
			return
		}

		if requesterRole == models.RoleManager && *body.Role == models.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "manager cannot assign admin role",
			})
			return
		}

		user.Role = *body.Role
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}