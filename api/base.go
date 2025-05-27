package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/response"
)

type BaseApi struct {
}

var store = base64Captcha.DefaultMemStore

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
func (b *BaseApi) GetEmailVerificationCode(c *gin.Context) {
	// emailConfig := global.Config.Email

}

// func (b *BaseApi) GetQQLink(c *gin.Context) {
// 	qq := global.Config.QQ
// }
