package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid claims"})
			return
		}

		if id, exists := claims["id"]; exists {
			c.Set("userID", fmt.Sprintf("%v", id))
		}
		if email, exists := claims["email"]; exists {
			c.Set("email", fmt.Sprintf("%v", email))
		}
		if role, exists := claims["role"]; exists {
			c.Set("role", fmt.Sprintf("%v", role))
		}
		if slug, exists := claims["slug"]; exists {
			c.Set("slug", fmt.Sprintf("%v", slug))
		}
		c.Next()
	}
}
