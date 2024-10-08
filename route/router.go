package route

import (
	"c2n/api/handler"
	"c2n/middleware"

	_ "c2n/docs" // 替换为你的项目路径

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/hello", helloHandler.GetHelloWord)
	r.POST("/encode/sign_registration", encodeHandler.SignRegistration)
	r.POST("/encode/sign_participation", encodeHandler.SignParticipation)
	r.GET("/product/base_info", productContractHandler.BaseInfo)
	r.GET("/product/list", productContractHandler.List)
	r.POST("/update", productContractHandler.Update)
	return r
}
