package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	error2 "github.com/supkit/tabe/error"
	"net/http"
	"strings"
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
		contentType := ctx.Request.Header.Get("Content-Type")

		// bind json
		if strings.Contains(contentType, "application/json") {
			if err := ctx.ShouldBindJSON(&req); err != nil {
				rsp = errorMessage(err, err.Error())
				ctx.JSON(http.StatusOK, rsp)
				return
			}
		}

		// bind queryString
		if strings.Contains(contentType, "text/html") || len(contentType) == 0 {
			if err := ctx.ShouldBindQuery(&req); err != nil {
				rsp = errorMessage(err, err.Error())
				ctx.JSON(http.StatusOK, rsp)
				return
			}
		}

		// bind formData => multipart/form-data || application/x-www-form-urlencoded
		if strings.Contains(contentType, "multipart/form-data") ||
			strings.Contains(contentType, "application/x-www-form-urlencoded") {
			if err := ctx.ShouldBind(&req); err != nil {
				rsp = errorMessage(err, err.Error())
				ctx.JSON(http.StatusOK, rsp)
				return
			}
		}

		data, err := handler(ctx, req)
		rid, ok := ctx.Value("rid").(string)
		if !ok {
			rid = ""
		}
		rsp.ID = rid

		if err != nil {
			rsp = errorMessage(err, err.Error())
			ctx.JSON(http.StatusOK, rsp)
			return
		}

		rsp.Code = 0
		rsp.Message = "success"
		rsp.Data = data
		ctx.JSON(http.StatusOK, rsp)
	}
}

func errorMessage(err error, message string) (rsp ResponseData) {
	rsp = ResponseData{}
	var err2 error2.Error
	if len(message) == 0 {
		message = "unknown error"
	}

	if errors.As(err, &err2) {
		rsp.Code = err2.Code()
		rsp.Message = err2.Message()
		rsp.Data = nil
	} else {
		rsp.Code = 1
		rsp.Message = message
		rsp.Data = nil
	}

	return
}
