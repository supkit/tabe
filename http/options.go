package http

import "github.com/gin-gonic/gin"

// Option set
type Options struct {
	Addr   string
	Mode   string
	Router func(engine *gin.Engine)
}

// Option 调用参数工具函数
type Option func(*Options)

// WithMode set mode
func WithMode(mode string) Option {
	return func(o *Options) {
		o.Mode = mode
	}
}

// WithAddr set address
func WithAddr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}

// WithRouter set router
func WithRouter(router func(engine *gin.Engine)) Option {
	return func(o *Options) {
		o.Router = router
	}
}
