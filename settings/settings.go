package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(MultipleConfig)

type MultipleConfig struct {
	*AppConfig   `mapstructure:"app"`
	*MysqlConfig `mapstructure:"mysql"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_Idle_conns"`
}

func InitConfiguration() (err error) {
	viper.SetConfigFile("conf/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Init config file failed, err: %v\n", err))
		return
	}

	//反序列化到结构体
	if err := viper.Unmarshal(Config); err != nil {
		fmt.Printf("viper unmarshal failed, err: %v\n", err)
	}

	// 监控并重新读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file is changed...", in.Name)
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Printf("viper unmarshal failed, err: %v\n", err)
		}
	})
	return
}
