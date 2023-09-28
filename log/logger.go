package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// logger zap logger
var logger *zap.Logger

// defaultMaxAge 默认日志保留时间
var defaultMaxAge = time.Hour * 7 * 24

// defaultCallerSkip 默认跳过调用者数量
var defaultCallerSkip = 1

// defaultFileFormat 默认日志文件格式
var defaultFileFormat = "%Y-%m-%d"

// New init
func New(filePath string, opt ...Option) *zap.Logger {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := getEncoder()
	options := Options{
		MaxAge:        defaultMaxAge,
		FileFormat:    defaultFileFormat,
		DistWarnLevel: false,
		CallerSkip:    defaultCallerSkip,
	}

	for _, o := range opt {
		o(&options)
	}

	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.WarnLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := NewWriter(filePath, options)
	if options.DistWarnLevel {
		filePath = filePath + "-error"
	}

	warnWriter := NewWriter(filePath, options)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(options.CallerSkip))
	return logger
}

// Debug debug level
func Debug(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	logger.Debug(msg)
}

// Info info level
func Info(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	logger.Info(msg)
}

// Warn warn level
func Warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	logger.Warn(msg)
}

// Error error level
func Error(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	logger.Error(msg)
}

// Fatal fatal level
func Fatal(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	logger.Fatal(msg)
}
