package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/setting"
	"go-web/routers/api"
	v1 "go-web/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/sys/login", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 登录
		apiv1.POST("/sys/login")

		// 获取部门列表
		apiv1.GET("/company/department", v1.GetDeptList)
		//获取部门详情
		apiv1.GET("/company/department/:id", v1.GetDeptDetail)
		//获取部门负责人列表
		apiv1.GET("/sys/user/simple", v1.GetDeptManagerList)
		//修改部门详情
		apiv1.PUT("/company/department/:id", v1.ModifyDept)
		//删除部门
		apiv1.DELETE("/company/department/:id", v1.DelDept)
		// 新增部门
		apiv1.POST("/company/department", v1.AddDept)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin",
		})
	})

	return r
}
