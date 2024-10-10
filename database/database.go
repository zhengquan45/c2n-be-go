package database

import (
	"c2n/config"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitializeDB() *gorm.DB {
	// 读取数据库配置
	dbUser := config.AppConfig.Database.User         // 数据库用户名
	dbPassword := config.AppConfig.Database.Password // 数据库密码
	dbHost := config.AppConfig.Database.Host         // 数据库主机地址
	dbPort := config.AppConfig.Database.Port         // 数据库端口
	dbName := config.AppConfig.Database.Name         // 数据库名称

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	gormLogger := &LogrusLogger{Log: log.StandardLogger()}

	//配置 GORM 连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数
		},
		Logger: gormLogger,
	})
	if err != nil {
		log.Error("failed to connect database:", err)
	}

	// 配置数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("Failed to configure database pool: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)           // 空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间

	DB = db

	log.Info("Database connection established")
	return DB
}
