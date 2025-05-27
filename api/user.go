package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/request"
	"swiftDaily_myself/model/response"
)

type UserApi struct {
}

func (u *UserApi) Login(c *gin.Context) {
	var req request.Login
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(req, c)
}
