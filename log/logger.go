package log

import (
	"context"
	"fmt"
	"github.com/segmentio/ksuid"
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

// defaultRid requestId
var defaultRid = ksuid.New().String()

// ridKey key
var ridKey = "rid"

// New init
func New(filePath string, opt ...Option) *zap.Logger {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := getEncoder()
	options := Options{
		MaxAge:        defaultMaxAge,
		FileFormat:    defaultFileFormat,
		DistWarnLevel: false,
		CallerSkip:    defaultCallerSkip,
		Hooks:         []zapcore.WriteSyncer{},
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

	infoSyncs := []zapcore.WriteSyncer{zapcore.AddSync(infoWriter)}
	warnSyncs := []zapcore.WriteSyncer{zapcore.AddSync(warnWriter)}

	for _, hook := range options.Hooks {
		infoSyncs = append(infoSyncs, hook)
		warnSyncs = append(warnSyncs, hook)
	}

	infoMultiWriter := zapcore.NewMultiWriteSyncer(infoSyncs...)
	warnMultiWriter := zapcore.NewMultiWriteSyncer(warnSyncs...)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoMultiWriter, infoLevel),
		zapcore.NewCore(encoder, warnMultiWriter, warnLevel),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(options.CallerSkip))
	return logger
}

// With context
func With(ctx context.Context) *zap.Logger {
	rid := ctx.Value(ridKey)
	if rid == nil {
		rid = defaultRid
	}
	return logger.With(zap.String(ridKey, rid.(string)))
}

// Debug debug level
func Debug(ctx context.Context, format string, args ...any) {
	With(ctx).Debug(fmt.Sprintf(format, args...))
}

// Info info level
func Info(ctx context.Context, format string, args ...any) {
	With(ctx).Info(fmt.Sprintf(format, args...))
}

// Warn warn level
func Warn(ctx context.Context, format string, args ...any) {
	With(ctx).Warn(fmt.Sprintf(format, args...))
}

// Error error level
func Error(ctx context.Context, format string, args ...any) {
	With(ctx).Error(fmt.Sprintf(format, args...))
}

// Fatal fatal level
func Fatal(ctx context.Context, format string, args ...any) {
	With(ctx).Fatal(fmt.Sprintf(format, args...))
}
