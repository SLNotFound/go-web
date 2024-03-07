package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web/config"
)

func main() {
	// 初始化配置
	if err := config.InitConfiguration(); err != nil {
		fmt.Printf("init configuration failed, err: %v\n", err)
		return
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin")
	})
	router.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
}
