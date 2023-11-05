package error

import (
	"fmt"
	"testing"
)

var errors Error

var ErrFoo = errors.Errorf(1000, "foo = %s")

func TestError_New(t *testing.T) {
	err := InvalidParameter("bar test error message")
	if err, ok := err.(Error); ok {
		fmt.Printf("code=%d message=%s\n", err.Code(), err.Message())
	} else {
		fmt.Printf("error=%v\n", err.Error())
	}

	err = ErrFoo("custom foo error")
	if err, ok := err.(Error); ok {
		fmt.Printf("foo code=%d message=%s\n", err.Code(), err.Message())
	} else {
		fmt.Printf("foo error=%v\n", err.Error())
	}
}

func InvalidParameter(params ...any) (err error) {
	err = errors.New(10000, "request params error params is = %v", params...)
	return
}
