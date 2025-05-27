package request

type Login struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captch" binding:"required"`
	CaptchaID string `json:"captcha_id" binding:"required"`
}
