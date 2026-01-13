package middlewares

import (
    "strings"
    "net/http"
    "demowebgo/utlis" 
    "github.com/gin-gonic/gin"
    "log"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Yêu cầu cung cấp mã xác thực"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Định dạng xác thực không hợp lệ"})
            c.Abort()
            return
        }

        claims, err := utils.ValidateAccessToken(parts[1])
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ hoặc đã hết hạn"})
            c.Abort()
            return
        }

        c.Set("user_id", claims.UserID)
        c.Set("email", claims.Email)
        c.Set("role", claims.Role)
        log.Printf("[Middleware] Đã xác thực UserID: %v", claims.UserID)
        c.Next()
    }
}
