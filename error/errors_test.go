package error

import (
	"fmt"
	"testing"
)

var errors Error

var ErrFoo = errors.Errorf(1000, "foo = %s")

func TestError_New(t *testing.T) {
	err := InvalidParameter("bar test error message")
	fmt.Printf("err is %v\n", err)
	err = ErrFoo("test")
	fmt.Println(err)
}

func InvalidParameter(params ...any) (err error) {
	err = errors.New(10000, "request params error params is = %v", params...)
	return
}
