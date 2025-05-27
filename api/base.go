package api

import (
	"github.com/gin-gonic/gin"
	"swiftDaily_myself/global"
)

type BaseApi struct {
}

func (b *BaseApi) GetCaptch(c *gin.Context) {
	captcha := global.Config.Captcha

}
func (b *BaseApi) GetEmailVerificationCode(c *gin.Context) {
	email := global.Config.Email
}

// func (b *BaseApi) GetQQLink(c *gin.Context) {
// 	qq := global.Config.QQ
// }
