package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Run http server run
func Run(addr string, router func(engine *gin.Engine)) {
	engine := gin.New()
	router(engine)
	err := engine.Run(addr)
	if err != nil {
		fmt.Printf("gin http server run error: %v\n", err)
	}
}
