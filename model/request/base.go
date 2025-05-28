package request

type SendEmailVerificationCode struct {
	Email     string `json:"email" binding:"required,email"` // 这个email是要存入store的
	CaptchaID string `json:"captcha_id" binding:"required"`
	Captcha   string `json:"captcha" binding:"required,len=6"`
}
