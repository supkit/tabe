package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRun(t *testing.T) {
	Run(":8089", func(engine *gin.Engine) {
		engine.GET("/", func(ctx *gin.Context) {
			fmt.Println("http server run")
		})
	})
}
