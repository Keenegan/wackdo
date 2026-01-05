package controllers_user

import (
	"time"
	"wackdo/src/initializers"
	"wackdo/src/models"
	"wackdo/src/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.BindJSON(&body)

	var user models.User
	if err := initializers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.Error(&service.InvalidParamError{
			Reason: "invalid credentials",
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.Error(&service.InvalidParamError{
			Reason: "invalid credentials",
		})
		return
	}

	claims := initializers.JwtClaims{
		UserID: user.ID,
		Role:   string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(initializers.JwtSecret)

	c.JSON(200, gin.H{
		"token": signedToken,
		"role":  user.Role,
	})
}

