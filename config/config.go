package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(MultipleConfig)

type MultipleConfig struct {
	*AppConfig `mapstructure:"app"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

type MysqlConfig struct {
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func InitConfiguration() (err error) {
	viper.SetConfigFile("config/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("init config file failed, err: %v\n", err))
		return
	}

	// 反序列化到结构体
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper unmarshal failed, err: %v\n", err)
	}

	// 监控配置文件 并重新读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file is changed...", in.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper unmarshal failed, err: %v\n", err)
		}
	})
	return nil
}
