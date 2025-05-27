package config

type Zap struct {
	Level          string `json:"level" yaml:"level"`
	Filename       string `json:"filename" yaml:"filename"`
	MaxSize        int    `json:"max_size" yaml:"max_size" mapstructure:"max_size"`
	MaxAge         int    `json:"max_age" yaml:"max_age" mapstructure:"max_age"`
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print" mapstructure:"is_console_print"`
	MaxBackups     int    `json:"max_backs" yaml:"max_backs" mapstruct:"max_backs"`
}
