package log

import (
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

// Options 参数
type Options struct {
	// MaxAge 日志最大保留时间(天)
	MaxAge time.Duration
	// FileFormat 未见时间格式
	FileFormat string
}

// Option 调用参数工具函数
type Option func(*Options)

// WithMaxAge 设置最大保留天数
func WithMaxAge(n time.Duration) Option {
	return func(o *Options) {
		o.MaxAge = n
	}
}

func getWriter(filename string, opt Options) io.Writer {
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotate.New(
		filename+".%Y-%m-%d",
		rotate.WithMaxAge(time.Hour*24*7),
		rotate.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}

	return hook
}
