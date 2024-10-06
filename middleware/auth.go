package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" || !validateToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// validateToken 验证 Token 的逻辑
func validateToken(token string) bool {
	// 这里是 Token 验证的逻辑，实际项目中可以调用 JWT 验证等
	return token == "valid-token" // 示例：简单验证
}
