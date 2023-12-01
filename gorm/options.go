package gorm

import "time"

// Options 可设置的参数
type Options struct {
	DSN               string
	DefaultStringSize uint
	MaxIdleConn       int
	MaxOpenConn       int
	ConnMaxLifetime   time.Duration
}

// Option 调用参数工具函数
type Option func(*Options)

// WithDSN dsn
func WithDSN(dsn string) Option {
	return func(o *Options) {
		o.DSN = dsn
	}
}

// WithStringSize 字符串长度
func WithStringSize(size uint) Option {
	return func(o *Options) {
		o.DefaultStringSize = size
	}
}

// WithMaxIdleConn 空闲连接数
func WithMaxIdleConn(n int) Option {
	return func(o *Options) {
		o.MaxIdleConn = n
	}
}

// WithMaxOpenConn 并发数
func WithMaxOpenConn(n int) Option {
	return func(o *Options) {
		o.MaxOpenConn = n
	}
}

// ConnMaxLifetime 可重用连接得最大时间长度
func ConnMaxLifetime(n time.Duration) Option {
	return func(o *Options) {
		o.ConnMaxLifetime = n
	}
}
