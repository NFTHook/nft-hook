package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type appConfig struct {
	Http struct {
		Port string
	}
	Mysql struct {
		Host     string
		Port     string
		Username string
		Password string
		Db       string
	}
	API struct {
		Key string
	}
	Log struct {
		Level string
	}
	AES struct {
		KeyPath string
	}
	Key struct {
		VegaKey []byte
	}
	Chains struct {
		Ethereum  string
		BSC       string
		BscTestWs string
	}
	Google struct {
		ClientId string
	}
	Openai struct {
		Apikey string `yaml:"apikey"`
	}
	Env struct {
		WebUrl     string
		Log        string
		En_i18n    string
		Zh_cn_i18n string
	}
	WxPay struct {
		CallbackUrl string
	}
}

var appCfg = &appConfig{}

func init() {

	env := os.Getenv("NFTHOOK_APP_ENV")
	if env == "production" {
		viper.SetConfigFile("./config/pro.yml")
	} else {
		viper.SetConfigFile("./config/dev.yml")
	}

	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("无法读取配置文件：%s", err)
	}

	err = viper.Unmarshal(appCfg)
	if err != nil {
		log.Fatalf("无法解析配置：%s", err)
	}
}

func Get() *appConfig {
	return appCfg
}
