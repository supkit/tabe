package http

import "github.com/gin-gonic/gin"

// Options set
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

// WithRouter set router
func WithRouter(router func(engine *gin.Engine)) Option {
	return func(o *Options) {
		o.Router = router
	}
}
