package error

import (
	"fmt"
	"regexp"
	"strconv"
)

// Error def error
type Error struct {
	code    uint32
	message string
}

// errMessage error message format
var errMessage = "tabe framework error code:%d message:%s"

// SystemErrCode system error code
var SystemErrCode = uint32(900001)

// Error interface instance of error
func (e Error) Error() string {
	return fmt.Sprintf(errMessage, e.code, e.message)
}

// New create error
func (e Error) New(code uint32, message string) Error {
	return Error{
		code:    code,
		message: message,
	}
}

// getCode get code
func (e Error) getCode() uint32 {
	return e.code
}

// getMessage get message
func (e Error) getMessage() string {
	return e.message
}

// Parse parse error message
func (e Error) Parse(error string) (code uint32, message string) {
	reg, err := regexp.Compile(`framework error code:([0-9]+) message:(.+)`)
	if err != nil {
		return SystemErrCode, err.Error()
	}

	match := reg.FindStringSubmatch(error)
	if len(match) == 0 {
		return SystemErrCode, error
	}

	c, err := strconv.Atoi(match[1])
	if err != nil {
		return SystemErrCode, err.Error()
	}

	if len(match) < 3 {
		return SystemErrCode, "match error message fail"
	}

	return uint32(c), match[2]
}
