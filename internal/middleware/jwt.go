package middleware

import (
	"backend-go-demo/internal/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func parseToken(tokenStr string) (*JWTData, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
		return config.GetJWTKey()
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTData); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		jwtData, err := parseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// this JWT is stateless and this is only stored during the request...
		c.Set(config.ClaimUserID, jwtData.UserID)
		c.Set(config.ClaimUsername, jwtData.Username)
		c.Next()
	}
}
