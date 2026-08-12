package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireRoles validates that the authenticated user has at least one of the allowed roles.
func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	allowed := make(map[string]bool, len(allowedRoles))
	for _, role := range allowedRoles {
		allowed[strings.ToLower(strings.TrimSpace(role))] = true
	}

	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "role claim missing"})
			return
		}

		role, ok := roleValue.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "invalid role claim"})
			return
		}

		if !allowed[strings.ToLower(strings.TrimSpace(role))] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden: insufficient role"})
			return
		}

		c.Next()
	}
}
