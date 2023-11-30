package gorm

// Option 调用参数工具函数
type Option func(*Options)

// WithDSN dsn
func WithDSN(dsn string) Option {
	return func(o *Options) {
		o.DSN = dsn
	}
}
