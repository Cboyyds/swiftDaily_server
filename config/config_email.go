package config

type Email struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	From     string `json:"from" yaml:"from"`
	NickName string `json:"nickname" yaml:"nickname"`
	Secret   string `json:"secret" yaml:"secret"`
	IsSSL    bool   `json:"is_ssl" yaml:"is_ssl" mapstructure:"is_ssl"`
}
