package response

type Captcha struct {
	PicPath   string `json:"pic_path"`
	CaptchaID string `json:"captcha_id"`
}
