package config

import (
	"gorm.io/gorm/logger"
	"strconv"
	"strings"
)

type Mysql struct {
	Host         string `json:"host" yaml:"host"`
	Port         int    `json:"port" yaml:"port"`
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	DbName       string `json:"db_name" yaml:"db_name" mapstructure:"db_name"`
	Config       string `json:"config" yaml:"config"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns" mapstructure:"max_open_conns"`
	LogMode      string `json:"log_mode" yaml:"log_mode" mapstructure:"log_mode"` // 为日志的输出模式，如info，slient等
}

func (m *Mysql) DSN() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")" + m.DbName + "?" + m.Config
}
func (m Mysql) LogLevel() logger.LogLevel {
	switch strings.ToLower(m.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
	
}
