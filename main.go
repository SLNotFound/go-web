package main

import (
	"fmt"
	"go-web/dao/mysql"
	"go-web/routers"
	"go-web/settings"
)

func main() {
	// 1、加载配置
	if err := settings.InitConfiguration(); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}

	// 2、初始化mysql连接
	if err := mysql.InitDB(settings.Config.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	defer mysql.Close()

	r := routers.SetUp(settings.Config.Mode)
}
