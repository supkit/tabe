package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNew(t *testing.T) {
	router := func(e *gin.Engine) {
		e.GET("/", func(context *gin.Context) {
			fmt.Println("http server run")
		})
	}

	opt := []Option{
		WithMode(gin.DebugMode),
		WithRouter(router),
	}

	New(":8082", opt...)
}
