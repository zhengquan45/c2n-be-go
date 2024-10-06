package config

import (
	"c2n/logger"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Env      string `mapstructure:"ENV"`
	Port     string `mapstructure:"PORT"`
	Database struct {
		Host     string `mapstructure:"HOST"`
		Port     string `mapstructure:"PORT"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		Name     string `mapstructure:"NAME"`
	} `mapstructure:"DATABASE"`
}

var AppConfig Config

// LoadConfig 加载配置文件
func LoadConfig() {
	env := os.Getenv("GO_ENV")
	if env != "" {
		viper.SetConfigName("config." + env) // 配置文件名（不带扩展名）
	} else {
		viper.SetConfigName("config") // 配置
	}
	viper.SetConfigType("yml")      // 配置文件类型
	viper.AddConfigPath("./config") // 添加配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Errorf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		logger.Log.Errorf("Unable to decode into struct, %v", err)
	}
}
