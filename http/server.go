package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// New create http server
func New(address string, opt ...Option) {
	options := Options{
		Mode: gin.DebugMode,
		Addr: address,
	}

	for _, o := range opt {
		o(&options)
	}

	engine := gin.New()
	gin.SetMode(options.Mode)

	// register router
	options.Router(engine)

	err := engine.Run(options.Addr)

	if err != nil {
		fmt.Printf("gin http server run error: %v\n", err)
	}
}
