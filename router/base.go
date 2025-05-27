package router

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	// 1.获取验证码
	// 2，获取邮箱验证码验证
	// 3，邮箱注册
	// 4，邮箱登录
}
