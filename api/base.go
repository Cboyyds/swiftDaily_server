package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/request"
	"swiftDaily_myself/model/response"
)

type BaseApi struct {
}

var store = base64Captcha.DefaultMemStore // 这个可以用来判断获取的验证码是不是刚才那个想登录注册的邮箱发过来的

func (b *BaseApi) GetCaptch(c *gin.Context) {
	captchaConfig := global.Config.Captcha
	// 创建base64编码验证码驱动器
	driver := base64Captcha.NewDriverDigit(
		captchaConfig.Height,
		captchaConfig.Width,
		captchaConfig.Length,
		captchaConfig.MaxSkew,
		captchaConfig.DotCount,
	)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.Log.Error("Failed to generate captcha", zap.Error(err))
		response.FailWithMessage("Failed to generate captcha", c)
		return
	}
	response.OKWithData(response.Captcha{
		CaptchaID: id,
		PicPath:   b64s,
	}, c)

}

func (b *BaseApi) SendEmailVerificationCode(c *gin.Context) {
	var req request.SendEmailVerificationCode
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error("Failed to bind JSON", zap.Error(err))
		response.FailWithMessage("Failed to bind JSON", c)
		return
	}
	// 验证码正确则发送邮件
	if store.Verify(req.CaptchaID, req.Captcha, true) {
		if err := baseService.SendEmailVerificationCode(c, req.Email); err != nil {
			response.FailWithMessage("Failed to send email", c)
			global.Log.Error("Failed to send email", zap.Error(err))
			return
		}
		response.OK(c)
		return
	}
	response.FailWithMessage("Failed to verify captcha", c)
	return
}

// func (b *BaseApi) GetQQLink(c *gin.Context) {
// 	qq := global.Config.QQ
// }
