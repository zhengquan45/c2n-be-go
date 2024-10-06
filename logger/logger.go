package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// InitLogger 初始化日志记录器
func InitLogger() {
	env := os.Getenv("GO_ENV")
	Log = logrus.New()

	// 设置日志格式为文本（也可以设置为 JSON）
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	// 设置输出到文件和控制台
	file, err := os.OpenFile("./output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// log.Fatal("Failed to log to file, using default stderr")
	}
	// 根据环境来设置日志输出
	if env == "production" {
		// 生产环境只写入文件
		Log.SetLevel(logrus.InfoLevel)
		Log.SetOutput(file)
	} else {
		Log.SetLevel(logrus.DebugLevel)
		// 开发环境，写入文件和控制台
		Log.SetOutput(io.MultiWriter(os.Stdout, file))
	}
}
