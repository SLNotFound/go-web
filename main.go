package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-web/config"
	"go-web/router"
	"net/http"
)

func main() {
	// 初始化配置
	if err := config.InitConfiguration(); err != nil {
		fmt.Printf("init configuration failed, err: %v\n", err)
		return
	}

	// 初始化路由
	r := router.Setup(config.Conf.Mode)

	// 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	srv.ListenAndServe()
}
