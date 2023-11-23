package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/service/login_service"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.GET("/auth", api.GetAuth)
	r.POST("/sys/login", login_service.UserLogin)
	return r
}
