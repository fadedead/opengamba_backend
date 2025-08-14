package middleware

import (
	"net/http"
	"open_gamba/internal/auth"
	"open_gamba/internal/user"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorization")
		if err != nil || token != "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authorization token invalid",
			})
			c.Abort()
			return
		}

		claims, err := auth.ValidateToken(string(token))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("roles", claims.Roles)

		c.Next()
	}
}

func RequireAnyRole(allowedRoles ...user.AccessRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user roles from context (set by JWT middleware)
		rolesInterface, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authentication required",
			})
			c.Abort()
			return
		}
		userRoles, ok := rolesInterface.([]user.AccessRole)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid role format in token",
			})
			c.Abort()
			return
		}
		// Check if user has any of the allowed roles
		if !hasAnyRole(userRoles, allowedRoles) {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Insufficient permissions for this endpoint",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func hasAnyRole(userRoles []user.AccessRole, allowedRoles []user.AccessRole) bool {
	for _, userRole := range userRoles {
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				return true
			}
		}
	}
	return false
}
