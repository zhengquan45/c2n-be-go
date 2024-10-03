package routes

import (
	"c2n/controllers"
	"c2n/middlewares"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 获取当前环境
	env := os.Getenv("GO_ENV")
	r := gin.Default()
	r.Use(middlewares.CustomRecovery())

	// 在生产环境中将日志写入文件
	if env == "production" {
		// 创建日志文件
		f, _ := os.Create("gin.log")

		// 设置 Gin 的日志输出到文件
		gin.DefaultWriter = io.MultiWriter(f)
		gin.DefaultErrorWriter = io.MultiWriter(f)

		// 仅保留 Recovery 中间件
		r.Use(gin.Recovery())
	} else {
		// 在开发环境中，启用 Logger 和 Recovery
		r.Use(gin.Logger(), gin.Recovery())
	}

	r.GET("/hello", controllers.GetHelloWord)
	r.POST("/encode/sign_registration", controllers.SignRegistration)
	return r
}
