package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	error2 "github.com/supkit/tabe/error"
	"net/http"
)

// HandlerFunc handler func
type HandlerFunc[T any] func(c *gin.Context, req T) (data any, err error)

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

// ResponseData response data
type ResponseData struct {
	ID      string `json:"id"`
	Code    uint32 `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Handler handler
func Handler[T any](handler HandlerFunc[T], req T) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rsp := ResponseData{}
		if err := ctx.BindJSON(&req); err != nil {
			fmt.Printf("debug bind json error: %v\n", err)
			err = nil
		}

		data, err := handler(ctx, req)
		rsp.ID = ctx.Value("rid").(string)
		if err != nil {
			var err error2.Error
			if errors.As(err, &err) {
				rsp.Code = err.Code()
				rsp.Message = err.Message()
			}
			rsp.Data = data
		} else {
			rsp.Code = 0
			rsp.Message = "success"
			rsp.Data = data
		}

		ctx.JSON(http.StatusOK, rsp)
	}
}
