package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestServer(t *testing.T) {
	Server(":8089", func(engine *gin.Engine) {
		engine.GET("/", func(ctx *gin.Context) {
			fmt.Println("http server run")
		})
	})
}
