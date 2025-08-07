package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	
	"github.com/company/go-enterprise-template/internal/config"
)

var Logger *logrus.Logger

// Init 初始化日志系统
func Init(cfg *config.LogConfig) error {
	Logger = logrus.New()

	// 设置日志级别
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return err
	}
	Logger.SetLevel(level)

	// 设置日志格式
	if cfg.Format == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	}

	// 设置输出
	var output io.Writer
	switch cfg.Output {
	case "file":
		// 使用lumberjack进行日志轮转
		output = &lumberjack.Logger{
			Filename:   filepath.Join("logs", "app.log"),
			MaxSize:    cfg.MaxSize,    // MB
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,     // days
			Compress:   cfg.Compress,
		}
	case "both":
		// 同时输出到文件和控制台
		fileOutput := &lumberjack.Logger{
			Filename:   filepath.Join("logs", "app.log"),
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}
		output = io.MultiWriter(os.Stdout, fileOutput)
	default:
		output = os.Stdout
	}

	Logger.SetOutput(output)

	// 确保logs目录存在
	if cfg.Output == "file" || cfg.Output == "both" {
		if err := os.MkdirAll("logs", 0755); err != nil {
			return err
		}
	}

	return nil
}

// GetLogger 获取日志实例
func GetLogger() *logrus.Logger {
	if Logger == nil {
		// 如果没有初始化，使用默认配置
		Logger = logrus.New()
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	}
	return Logger
}

// WithFields 创建带字段的日志条目
func WithFields(fields logrus.Fields) *logrus.Entry {
	return GetLogger().WithFields(fields)
}

// WithField 创建带单个字段的日志条目
func WithField(key string, value interface{}) *logrus.Entry {
	return GetLogger().WithField(key, value)
}

// WithError 创建带错误的日志条目
func WithError(err error) *logrus.Entry {
	return GetLogger().WithError(err)
}

// Debug 输出调试日志
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Debugf 输出格式化调试日志
func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

// Info 输出信息日志
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Infof 输出格式化信息日志
func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

// Warn 输出警告日志
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Warnf 输出格式化警告日志
func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

// Error 输出错误日志
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Errorf 输出格式化错误日志
func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

// Fatal 输出致命错误日志并退出程序
func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

// Fatalf 输出格式化致命错误日志并退出程序
func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

// Panic 输出panic日志并触发panic
func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}

// Panicf 输出格式化panic日志并触发panic
func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args...)
}