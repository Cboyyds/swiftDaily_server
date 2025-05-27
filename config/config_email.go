package config

type Email struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	From     string `json:"from" yaml:"from"`
	NickName string `json:"nick_name" yaml:"nick_name" mapstructure:"nick_name"`
	Secret   string `json:"secret" yaml:"secret"`
}
