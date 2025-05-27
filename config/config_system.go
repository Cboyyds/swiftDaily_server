package config

import "fmt"

type System struct {
	Host           string `json:"-" yaml:"host"`
	Port           int    `json:"-" yaml:"port"`
	Env            string `json:"-" yaml:"env"` // Gin 的环境类型，例如 "debug"、"release" 或 "test"
	RouterPrefix   string `json:"-" yaml:"router_prefix" mapstructure:"router_prefix"`
	OssType        string `json:"oss_type" yaml:"oss_type" mapstructure:"oss_type"`                      // 存储方式类型              //
	SessionsSecret string `json:"sessions_secret" yaml:"sessions_secret" mapstructure:"sessions_secret"` // 对应的对象存储服务类型，如 "local" 或 "qiniu",暂时先用到本地存储
}

func (s *System) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
