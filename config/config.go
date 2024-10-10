package config

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

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
	Owner struct {
		PrivateKey string `mapstructure:"PRIVATE_KEY"`
	} `mapstructure:"OWNER"`
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

	viper.SetConfigType("yml") // 配置文件类型

	configPath := findRoot() + "/config"
	viper.AddConfigPath(configPath) // 添加配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Errorf("Unable to decode into struct, %v", err)
	}
}

func findRoot() string {
	if os.Getenv("GO_ENV") == "production" {
		exePath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		return filepath.Dir(exePath)
	} else {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		// 不断向上查找 go.mod 文件
		for {
			if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
				return dir // 找到 go.mod 文件，返回当前目录
			}

			// 向上一级目录
			parent := filepath.Dir(dir)
			if parent == dir {
				// 已经到根目录，停止
				break
			}
			dir = parent
		}

		panic("could not find go.mod")
	}

}
