package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func New(filePath string, opt ...Option) *zap.Logger {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := getEncoder()
	options := Options{
		MaxAge: time.Hour * 7 * 24,
	}

	for _, o := range opt {
		o(&options)
	}

	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter(filePath, options)
	warnWriter := getWriter(filePath, options)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	logger = zap.New(core, zap.AddCaller())
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
