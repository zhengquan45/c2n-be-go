package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CustomRecovery 中间件，用于捕获 panic 并返回自定义的错误信息
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录错误信息
				log.Error("捕获到异常: %v\n", err)

				// 返回自定义的 JSON 错误响应
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  http.StatusInternalServerError,
					"message": fmt.Sprintf("内部服务器错误: %v", err),
				})

				// 阻止后续的处理
				c.Abort()
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
