package route

import (
	"c2n/api/handler"
	"c2n/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	encodeHandler *handler.EncodeHandler,
	helloHandler *handler.HelloHandler,
	productContractHandler *handler.ProductContractHandler,
) *gin.Engine {
	r := gin.New()
	// 使用自定义的 Recovery 中间件
	r.Use(middleware.CustomRecovery())
	// 使用自定义的 Logger 中间件
	// r.Use(middleware.CustomLogger())
	// 注册路由
	r.GET("/hello", helloHandler.GetHelloWord)
	r.POST("/encode/sign_registration", encodeHandler.SignRegistration)
	r.GET("/product_contract/base_info", productContractHandler.BaseInfo)
	return r
}
