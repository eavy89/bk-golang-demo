package middleware

import "github.com/golang-jwt/jwt/v5"

type JWTData struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
