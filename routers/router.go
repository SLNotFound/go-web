package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/sys/login", api.GetAuth)
	r.GET("/sys/profile", api.GetUserInfo)
	return r
}
