package router

import (
	"github.com/gin-gonic/gin"
	"swiftDaily_myself/api"
	"swiftDaily_myself/middleware"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	// userRouter := Router.Group("user")
	userPublicRouter := PublicRouter.Group("user")
	userLoginRouter := userPublicRouter.Use(middleware.LoginRecord()) // 这里可以加上一个中间件来记录登录的情况
	userApi := api.ApiGroupApp.UserApi
	{
		// userRouter.POST("logout", userApi.Login)
	}
	{
		userLoginRouter.POST("register", userApi.EmailRegister)
		userLoginRouter.POST("login", userApi.EmailLogin)
	}
}
