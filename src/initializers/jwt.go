package initializers

import "github.com/golang-jwt/jwt/v5"
import "os"

var JwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// todo rename package config ?
type JwtClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}