package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-web/models"
	"go-web/pkg/e"
	"go-web/pkg/logging"
	"go-web/pkg/util"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(20)"`
	Password string `valid:"Required; MaxSize(20)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist, _ := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": e.GetMsg(code),
		"success": true,
		"code":    code,
		"data":    data,
	})
}
