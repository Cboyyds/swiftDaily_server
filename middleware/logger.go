package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"swiftDaily_myself/global"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		
		cost := time.Since(start)
		global.Log.Info(path,
			// 记录响应状态码
			zap.Int("status", c.Writer.Status()),
			// 请求方法
			zap.String("method", c.Request.Method),
			// 请求路径
			zap.String("path", path),
			// 请求参数
			zap.String("query", query),
			// 请求耗时
			zap.String("cost", cost.String()),
			// Ip
			zap.String("ip", c.ClientIP()),
			// user-agent信息
			zap.String("user-agent", c.Request.UserAgent()),
			// 错误信息
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}

// GinRecovery 用于捕获和处理请求中的panic错误
// 该错误确保服务在遇到未处理的异常时不会崩溃，并通过日志系统提供详细的错误跟踪
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// 检查是否发生了panic错误
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connect reset by peer") {
							brokenPipe = true
						}
					}
				}
				// 获取请求信息，包括请求体等
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				// 如果时broken pipe'，则只记录错误信息，不记录堆栈信息
				if brokenPipe {
					global.Log.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// 由于链接断开，不能再向客户端写入状态码
					_ = c.Error(err.(error))
					c.Abort()
					return
				}
				
				// 如果是其他类型的panic，根据stack参数决定是否记录堆栈信息
				if stack {
					// 记录详细的错误和堆栈信息
					global.Log.Error("[Recovery from panic ]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)))
				}
			}
			
		}()
		// 继续执行后续请求处理
		c.Next()
	}
}
