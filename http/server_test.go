package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNew(t *testing.T) {
	router := func(e *gin.Engine) {
		e.POST("/user", Handler(User, UserReq{}))
		e.POST("/post", Handler(func(c *gin.Context, req FormDataReq) (data any, err error) {
			data = fmt.Sprintf("%+v", req)
			return
		}, FormDataReq{}))
		e.GET("/query", Handler(func(c *gin.Context, req any) (data any, err error) {
			data = fmt.Sprintf("%+v", req)
			return
		}, nil))
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
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type FormDataReq struct {
	Key  string `form:"key" binding:"required"`
	Name string `form:"name" binding:"required,email"`
}

type QueryDataReq struct {
	Key  string `form:"key" binding:"required"`
	Name string `form:"name" binding:"required,email"`
}

func User(ctx *gin.Context, req UserReq) (data any, err error) {
	data = req.Name
	return
}
