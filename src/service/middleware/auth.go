package middleware

import (
	"net/http"
	"strings"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(allowedRoles ...models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.ParseWithClaims(
			tokenString,
			&initializers.JwtClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return initializers.JwtSecret, nil
			},
		)

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(*initializers.JwtClaims)
		userRole := models.Role(claims.Role)

		// Admin has all rights, so they bypass authentication middleware
		if userRole != models.RoleAdmin {
			for _, role := range allowedRoles {
				if userRole == role {
					c.Set("userID", claims.UserID)
					c.Set("role", userRole)
					c.Next()
					return
				}
			}

			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
