package config

import (
	"log"

	"github.com/spf13/viper"
)

var App struct {
	Bot      BotConf   `yaml:"bot"`
	Log      LogConf   `yaml:"log"`
	Spark    SparkConf `yaml:"spark"`
	Database Database  `yaml:"database"`
}

type Database struct {
	Sqlite Sqlite `yaml:"sqlite"`
}

type Sqlite struct {
	Path string `yaml:"path"`
}

type SparkConf struct {
	Mode      string `yaml:"mode"`
	AppId     string `yaml:"appId"`
	ApiSecret string `yaml:"apiSecret"`
	ApiKey    string `yaml:"apiKey"`
}

type BotConf struct {
	Token string `yaml:"token"`
	Uin   uint64 `yaml:"uin"`
	AppId uint64 `yaml:"appId"`
}

type LogConf struct {
	Path string `yaml:"path"`
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read in config failed:", err)
	}
	viper.Unmarshal(&App)
}
