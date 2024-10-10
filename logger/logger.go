package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// InitLogger 初始化日志记录器
func InitLogger() {
	env := os.Getenv("GO_ENV")

	// 设置输出到文件和控制台
	file, err := os.OpenFile("./output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// log.Fatal("Failed to log to file, using default stderr")
	}
	// 根据环境来设置日志输出
	if env == "production" {
		// 生产环境只写入文件
		log.SetLevel(log.InfoLevel)
		log.SetOutput(file)
	} else {
		log.SetLevel(log.DebugLevel)
		// 开发环境，写入文件和控制台
		log.SetOutput(io.MultiWriter(os.Stdout, file))
	}
}
