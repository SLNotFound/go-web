package api

import (
	"github.com/gin-gonic/gin"
	"go-web/models"
	"go-web/pkg/e"
	"go-web/pkg/logging"
	"go-web/pkg/util"
	"net/http"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password""`
}

func GetAuth(c *gin.Context) {

	var auth Auth
	code := e.ERROR
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	data := make(map[string]interface{})

	isExist, err := models.CheckAuth(auth.Username, auth.Password)
	if isExist {
		token, err := util.GenerateToken(auth.Username, auth.Password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = e.SUCCESS
		}
	} else {
		logging.Info(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": e.GetMsg(code),
		"success": true,
		"code":    code,
		"data":    data,
	})
}
