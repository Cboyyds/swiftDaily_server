package config

type Jwt struct {
	AccessTokenSecret  string `json:"_" yaml:"access_token_secret" mapstructure:"access_token_secret"`
	RefreshTokenSecret string `json:"-" yaml:"refresh_token_secret" mapstructure:"refresh_token_secret"`
	AccessTokenExpire  string `json:"-" yaml:"access_token_expire" mapstructure:"access_token_expire"`
	RefreshTokenExpire string `json:"-" yaml:"refresh_token_expire" mapstructure:"refresh_token_expire"`
	Issuer             string `json:"issuer" yaml:"issuer"`
}
