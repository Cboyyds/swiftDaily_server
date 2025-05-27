package router

import (
	"github.com/gin-gonic/gin"
	"swiftDaily_myself/api"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	baseApi := api.ApiGroupApp.BaseApi
	{
		// 1.获取验证码
		baseRouter.POST("captcha", baseApi.GetCaptch)
		// 2.获取邮箱验证码验证
		baseRouter.POST("sendEmailVerificationCode", baseApi.GetEmailVerficationCode)
		// 3.获得qq登录链接
	}
}
