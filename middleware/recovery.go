package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// CustomRecovery 中间件，用于捕获 panic 并返回自定义的错误信息
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				// 记录错误信息
				log.Errorf("catch exception: %v\n", err)

				var response ErrorResponse
				// 检查错误类型
				switch e := err.(type) {
				case *BusinessError:
					// 处理业务异常
					response = NewErrorResponse(e.ReCode.GetCode(), e.ReCode.GetDesc())
					c.JSON(e.ReCode.GetCode(), response)
				default:
					// 处理系统异常
					response = NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
					c.JSON(http.StatusInternalServerError, response)
				}
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
