package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-web/models"
	"go-web/pkg/e"
	"go-web/pkg/util"
	"log"
	"net/http"
)

type auth struct {
	Mobile   string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	mobile := c.Query("mobile")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Mobile: mobile, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(mobile, password)
		if isExist {
			token, err := util.GenerateToken(mobile, password)
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
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    code,
		"message": "登录成功",
		"data":    data,
	})
}
