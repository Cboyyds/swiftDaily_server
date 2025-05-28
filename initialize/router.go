package initialize

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"swiftDaily_myself/global"
	"swiftDaily_myself/middleware"
	"swiftDaily_myself/router"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	gin.SetMode(global.Config.System.Env)
	// 配置日志记录中间件
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))                // 根据stack参数决定是否记录堆栈信息
	Router.StaticFS(global.Config.Upload.Path, http.Dir(global.Config.Upload.Path)) // http.Dir(...)：将本地目录转为 HTTP 文件系统接
	cookieStore := cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	Router.Use(sessions.Sessions("session", cookieStore))
	routerGroup := router.RouterGroupApp // 初始化各个路由和绑定对应的处理函数
	// 创建路由组
	publicGroup := Router.Group(global.Config.System.RouterPrefix)
	privateGroup := publicGroup
	privateGroup.Use(middleware.JWTAuth()) // 不要用privateGroup = publicGroup.Use(middleware.JWTAuth()),privateGroup会是IRouter类型的
	{
		routerGroup.InitBaseRouter(publicGroup) // 还未加入qq登录
	}
	{
		routerGroup.InitUserRouter(privateGroup, publicGroup)
	}
	return Router
}
