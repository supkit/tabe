package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	error2 "github.com/supkit/tabe/error"
	"net/http"
)

// HandlerFunc handler func
type HandlerFunc func(c *gin.Context) (data any, err error)

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
func Handler(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rsp := ResponseData{}
		data, err := handler(ctx)
		rsp.ID = ctx.Value("rid").(string)
		if err != nil {
			if err, ok := err.(error2.Error); ok {
				rsp.Code = err.Code()
				rsp.Message = err.Message()
			} else {
				rsp.Code = 10000
				rsp.Message = "system error"
			}
			rsp.Data = []string{}
		} else {
			rsp.Code = 0
			rsp.Message = "success"
			rsp.Data = data
		}

		ctx.JSON(http.StatusOK, rsp)
	}
}
