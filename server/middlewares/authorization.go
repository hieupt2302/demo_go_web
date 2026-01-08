package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}	
	}
	return false
}
func AuthorizeJWT(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_,existed:=c.Get("user_id");
	if !existed {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		c.Abort()
		return
	}
	role, existed := c.Get("role")
	if !existed {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
		c.Abort()
		return
	}
	if !contains(roles, role.(string)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		c.Abort()
		return
	}
		c.Next()
}
}
