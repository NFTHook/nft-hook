package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port         string       `mapstructure:"port"`
	SqliteConfig SQLiteConfig `mapstructure:"sqliteConfig"`
	Infura       InfuraConfig `mapstructure:"infura"`
}

type SQLiteConfig struct {
	FilePath string `mapstructure:"filePath"`
	Mode     string `mapstructure:"mode"`
}

type InfuraConfig struct {
	ProjectID     string `mapstructure:"projectId"`
	ProjectSecret string `mapstructure:"projectSecret"`
}

var AppCfg *AppConfig

func init() {
	// 使用Viper库读取配置文件
	viper.SetConfigFile("./config/dev.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %s", err))
	}

	// 使用Viper的Unmarshal方法将配置文件内容解析到AppCfg结构体中
	err = viper.Unmarshal(&AppCfg)
	if err != nil {
		panic(fmt.Errorf("error unmarshaling config: %s", err))
	}
}
