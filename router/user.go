package router

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userRouter := Router.Group("user")
	publicRouter := PublicRouter.Group("user")
	
}
