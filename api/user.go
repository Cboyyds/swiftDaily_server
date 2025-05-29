package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/database"
	"swiftDaily_myself/model/request"
	"swiftDaily_myself/model/response"
	"swiftDaily_myself/utils"
)

type UserApi struct {
}

func (u *UserApi) EmailRegister(c *gin.Context) {
	var req request.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Failed to bind JSON", c)
		global.Log.Error("Failed to bind JSON", zap.Error(err))
		return
	}
	// session := sessions.Default(c)
	// savedEmail := session.Get("email")
	// if savedEmail != nil || savedEmail.(string) == req.Email {
	// 	response.FailWithMessage("Email verification failed", c)
	// 	global.Log.Error("Email verification failed", zap.Error(errors.New("Email verification failed")))
	// 	return
	// }
	// savedCode := session.Get("verification_code")
	// if savedCode.(string) != req.VerificationCode {
	// 	response.FailWithMessage("verification code error", c)
	// 	global.Log.Error("verification code error", zap.Error(errors.New("verification code error")))
	// 	return
	// }
	// savedTime := session.Get("expire_time")
	// if savedTime.(int64) < time.Now().Unix() {
	// 	response.FailWithMessage("time expire", c)
	// 	global.Log.Error("time expire", zap.Error(errors.New("time expire")))
	// 	return
	// }
	var user database.User = database.User{
		Email:     req.Email,
		UserName:  req.Username,
		Password:  utils.BcryptHash(req.Password), // 对密码进行加密处理
		CompanyID: req.CompanyID,
		UUID:      uuid.Must(uuid.NewV6()),
	}
	if err := baseService.EmailRegister(user); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.Log.Error(err.Error(), zap.Error(err))
		return
	}
	response.OK(c)
	global.Log.Info("register success", zap.Any("user", user))
	return
}
func (u *UserApi) EmailLogin(c *gin.Context) {
	var req request.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Failed to bind JSON", c)
		global.Log.Error("Failed to bind JSON", zap.Error(err))
		return
	}
	//  判断验证码是否正确
	// if !store.Verify(req.CaptchaID, req.Captcha, true) {
	// 	response.FailWithMessage("验证码错误", c)
	// 	global.Log.Error("验证码错误", zap.Error(errors.New("验证码错误")))
	// 	return
	// }
	user := database.User{
		Email:    req.Account,
		Password: req.Password,
	}
	user, err := baseService.EmailLogin(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.Log.Error(err.Error(), zap.Error(err))
		return
	}
	response.OKWithDetail(user, "登录成功", c)
	global.Log.Info("login success", zap.Any("user", user))
}
