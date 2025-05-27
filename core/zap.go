package core

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"swiftDaily_myself/global"
)

// 该写入器利用 lumberjack 包，实现日志的滚动记录 ：轮转（Log Rotation）是一种日
// 志管理技术，当日志文件达到指定大小或时间限制时，系统会自动将当前日志文件归档，并创建一
// 个新的日志文件继续记录。github.com/natefinch/lumberjack包实现了这一功能，支持按文件
// 大小（MaxSize）、保留文件数量（MaxBackups）和保留天数（MaxAge）进行日志轮转。
func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger) // zapcore.Writer
}

// getEncoder 返回一个生产日志的编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func InitLogger() *zap.Logger {
	zapCfg := global.Config.Zap
	// 创建一个用于日志输出的writerSyncer
	writeSyncer := getLogWriter(zapCfg.Filename, zapCfg.MaxSize, zapCfg.MaxBackups, zapCfg.MaxAge)
	// 如果配置控制台输出，则添加
	if zapCfg.IsConsolePrint {
		writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
	}
	// 创建日志格式化的编辑器
	encoder := getEncoder()
	
	// 根据配置确定日志级别
	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(zapCfg.Level)); err != nil {
		log.Fatalf("Failed to parse log level: %v", err)
	}
	// 创建核心和实例
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}
