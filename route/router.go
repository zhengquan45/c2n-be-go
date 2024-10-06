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
	// 注册路由
	r.GET("/hello", helloHandler.GetHelloWord)
	r.POST("/encode/sign_registration", encodeHandler.SignRegistration)
	r.POST("/encode/sign_participation", encodeHandler.SignParticipation)
	r.GET("/product/base_info", productContractHandler.BaseInfo)
	r.GET("/product/list", productContractHandler.List)
	r.POST("/update", productContractHandler.Update)
	return r
}
