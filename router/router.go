package router

import (
	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	//gin.SetMode(mode)
	r := gin.New()
	gin.SetMode(mode)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, GIN")
	})

	r.POST("/sys/user")
	return r
}
