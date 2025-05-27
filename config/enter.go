package config

type Config struct {
	System  System  `json:"system" yaml:"system"`
	Mysql   Mysql   `json:"mysql" yaml:"mysql"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	Zap     Zap     `json:"zap" yaml:"zap"`
	Jwt     Jwt     `json:"jwt" yaml:"jwt"`
	Email   Email   `json:"email" yaml:"email"`
	Captcha Captcha `json:"captcha" yaml:"captcha"`
	Website Website `json:"website" yaml:"website"`
	Upload  Upload  `json:"upload" yaml:"upload"`
}
