package router

import (
	"github.com/gin-gonic/gin"
	"swiftDaily_myself/api"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userPublicRouter := PublicRouter.Group("user")
	userLoginRouter := userPublicRouter // 这里可以加上一个中间件来记录登录的情况
	userApi := api.ApiGroupApp.UserApi
	{
		userRouter.POST("logout", userApi.Login)
	}
	{
		userLoginRouter.POST("login", userApi.Login)
	}
}
