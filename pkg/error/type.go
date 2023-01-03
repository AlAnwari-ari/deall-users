package error

import (
	"errors"
	"fmt"
	"net/http"

	res "github.com/deall-users/pkg/response"
)

type Error struct {
	orig error
	msg  string
	code int
}

func (e *Error) Error() string {
	if e.orig != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.orig)
	}

	return e.msg
}

// Unwrap returns the wrapped error, if any.
func (e *Error) Unwrap() error {
	if e.orig != nil {
		return e.orig
	} else if e.msg != "" {
		return errors.New(e.msg)
	}

	return nil
}

// Code returns the code representing this error.
func (e *Error) Code() int {
	return e.code
}

func WrapErrorf(orig error, code int, format string, a ...interface{}) error {
	return &Error{
		code: code,
		orig: orig,
		msg:  fmt.Sprintf(format, a...),
	}
}

func NewErrorf(code int, format string, a ...interface{}) error {
	return WrapErrorf(nil, code, format, a...)
}

//	Err Resp representing this error.
func (e *Error) RespError(errData interface{}) *res.Response {
	return res.DefaultResponse(e.Error(), nil, errData, e.code)
}

func UnwrapErrorToResponse(err error) *res.Response {
	if errWrapped, ok := err.(*Error); ok {
		return errWrapped.RespError(nil)
	}

	return res.DefaultResponse(err.Error(), nil, nil, http.StatusInternalServerError)
}
