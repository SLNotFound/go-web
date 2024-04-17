package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-web/models"
	"go-web/pkg/e"
	"net/http"
)

// GetDeptList 获取部门列表
func GetDeptList(c *gin.Context) {
	maps := make(map[string]interface{})
	var data interface{}

	code := e.SUCCESS
	data = models.GetDepts(maps)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    code,
		"data":    data,
		"message": "获取成功",
	})
}

// GetDeptDetail 获取部门详情
func GetDeptDetail(c *gin.Context) {}

// GetDeptManagerList 获取部门负责人列表
func GetDeptManagerList(c *gin.Context) {}

// ModifyDept 修改部门详情
func ModifyDept(c *gin.Context) {}

// DelDept 删除部门
func DelDept(c *gin.Context) {

}

// AddDept 新增部门
func AddDept(c *gin.Context) {
	deptCode := c.Query("code")
	deptName := c.Query("name")
	managerId := c.GetInt("managerId")
	introduce := c.Query("introduce")
	deptPID := c.GetInt("pid")

	valid := validation.Validation{}
	valid.Required(deptCode, "code").Message("部门编码不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		models.AddDept(deptCode, deptName, introduce, deptPID, managerId)
	} else {
		code = e.ERROR_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    code,
		"data":    make(map[string]interface{}),
		"message": "添加成功",
	})
}
