package main

import (
	"fmt"
	"go-web/models"
	"go-web/pkg/logging"
	"go-web/pkg/setting"
	"go-web/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	router := routers.InitRouter()

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	router.Run(endPoint)

}
