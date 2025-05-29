package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ua-parser/uap-go/uaparser"
	"go.uber.org/zap"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/database"
	"swiftDaily_myself/service"
)

// LoginRecord 是一个中间件，用于记录登录日志
func LoginRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 异步记录日志
		go func() {
			gaodeService := service.ServiceGroupApp.GaodeService
			var userID uint
			var address string
			ip := c.ClientIP()
			loginMethod := c.DefaultQuery("flag", "email") // 若未传递flag参数，则默认为"email"
			userAgent := c.Request.UserAgent()

			// 从请求上下文中获取用户ID，确保获取到的是当前请求的正确用户ID
			if value, exists := c.Get("user_id"); exists {
				if id, ok := value.(uint); ok {
					userID = id
				}
			}
			// 获取用户IP的地理位置
			address = getAddressFromIP(ip, gaodeService)

			// 解析用户的浏览器、操作系统和设备信息
			os, deviceInfo, browserInfo := parseUserAgent(userAgent)

			// 创建登录记录
			login := database.Login{
				UserID:      userID,
				LoginMethod: loginMethod,
				IP:          "127",
				Address:     address,
				OS:          os,
				DeviceInfo:  deviceInfo,
				BrowserInfo: browserInfo,
			}
			// 将登录记录存储到数据库
			if err := global.DB.Create(&login).Error; err != nil {
				global.Log.Error("Failed to record login", zap.Error(err))
			}
		}()
	}
}

// 获取IP地址对应的地理位置信息
func getAddressFromIP(ip string, gaodeService service.GaodeService) string {
	res, err := gaodeService.GetLocationByIP(ip)
	if err != nil || res.Province == "" {
		return "未知"
	}
	if res.City != "" && res.Province != res.City {
		return res.Province + "-" + res.City
	}
	return res.Province
}

// 解析用户代理（User-Agent）字符串，提取操作系统、设备信息和浏览器信息
func parseUserAgent(userAgent string) (os, deviceInfo, browserInfo string) {
	// 解析用户代理（User-Agent）字符串，提取操作系统、设备信息和浏览器信息
	// 初始化返回值为原始User-Agent字符串，作为默认值
	os = userAgent
	deviceInfo = userAgent
	browserInfo = userAgent

	// 使用uaparser库解析User-Agent字符串
	parser := uaparser.NewFromSaved() // 创建一个预加载的解析器
	cli := parser.Parse(userAgent)    // 解析User-Agent字符串

	// 提取操作系统、设备信息和浏览器信息
	os = cli.Os.Family                 // 操作系统名称
	deviceInfo = cli.Device.Family     // 设备类型
	browserInfo = cli.UserAgent.Family // 浏览器名称

	return // 返回解析结果
}
