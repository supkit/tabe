package http

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNew(t *testing.T) {
	router := func(e *gin.Engine) {
		e.GET("/", Handler(User, UserReq{}))
	}

	opt := []Option{
		WithMode(gin.DebugMode),
		WithRouter(router),
	}

	New(":8082", opt...)
}

func TestHandler(t *testing.T) {
	Handler(User, UserReq{})
}

type UserReq struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func User(ctx *gin.Context, req UserReq) (data any, err error) {
	return
}
