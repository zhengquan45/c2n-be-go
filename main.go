package main

import (
	"c2n/api/handler"
	"c2n/config"
	"c2n/database"
	"c2n/logger"
	"c2n/repository"
	"c2n/route"
	"c2n/service"

	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 文件中的环境变量
	godotenv.Load()
	// 初始化日志记录器
	logger.InitLogger()
	// 加载配置文件
	config.LoadConfig()
	// 初始化数据库
	db := database.InitializeDB()
	// 实例化 repository
	productContractRepo := repository.NewProductContractRepository(db)

	// 实例化 service
	productContractService := service.NewProductContractService(productContractRepo)
	encodeService := service.NewEncodeService()

	// 实例化 handler
	productContractHandler := handler.NewProductContractHandler(productContractService)
	encodeHandler := handler.NewEncodeHandler(encodeService)
	helloHandler := handler.NewHelloHandler()

	r := route.SetupRouter(
		encodeHandler,
		helloHandler,
		productContractHandler,
	)

	r.Run(":9999") // 启动服务
}
