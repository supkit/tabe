package log

import (
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

// Options 参数
type Options struct {
	// MaxAge 日志最大保留时间(天)
	MaxAge time.Duration
	// FileFormat 未见时间格式
	FileFormat string
	// DistWarnLevel 区分warn以上的level
	DistWarnLevel bool
	// CallerSkip 跳过的调用者数量
	CallerSkip int
	Hooks      []zapcore.WriteSyncer
}

// Option 调用参数工具函数
type Option func(*Options)

// WithMaxAge 设置最大保留天数
func WithMaxAge(n time.Duration) Option {
	return func(o *Options) {
		o.MaxAge = n
	}
}

// WithFileFormat 设置文件格式
func WithFileFormat(format string) Option {
	return func(o *Options) {
		o.FileFormat = format
	}
}

// WithDistWarnLevel 设置区分warn以上的level的日志文件
func WithDistWarnLevel(dist bool) Option {
	return func(o *Options) {
		o.DistWarnLevel = dist
	}
}

// WithCallerSkip 设置跳过的调用者数量
func WithCallerSkip(skip int) Option {
	return func(o *Options) {
		o.CallerSkip = skip
	}
}

// WithHook 增加一个配置函数
func WithHooks(hooks ...zapcore.WriteSyncer) Option {
	return func(o *Options) {}
}

// NewWriter new writer
func NewWriter(filename string, opt Options) io.Writer {
	if len(opt.FileFormat) <= 0 {
		opt.FileFormat = ".%Y-%m-%d"
	}

	file := filename + opt.FileFormat

	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotate.New(
		file,
		rotate.WithMaxAge(time.Hour*24*7),
		rotate.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}

	return hook
}
