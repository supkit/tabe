package error

import (
	"fmt"
)

// Error def error
type Error struct {
	code    uint32
	message string
	arg     []any
}

// errMessage error message format
var errMessage = "system error code:%d message:%s"

// Error interface instance of error
func (e Error) Error() string {
	return fmt.Sprintf(errMessage, e.Code(), e.Message())
}

// Errorf error format
func (e Error) Errorf(code uint32, message string) func(arg ...any) error {
	return func(arg ...any) error {
		return e.New(code, message, arg...)
	}
}

// New create error
func (e Error) New(code uint32, message string, arg ...any) Error {
	return Error{
		code:    code,
		message: message,
		arg:     arg,
	}
}

// Code get code
func (e Error) Code() uint32 {
	return e.code
}

// Message get message
func (e Error) Message() string {
	if len(e.arg) > 0 {
		return fmt.Sprintf(e.message, e.arg...)
	}
	return e.message
}
