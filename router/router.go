package router

import "github.com/gin-gonic/gin"

func Setup(mode string) *gin.Engine {
	//gin.SetMode(mode)
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, GIN")
	})
	return r
}
