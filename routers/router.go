package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/auth", api.GetAuth)
	return r
}
