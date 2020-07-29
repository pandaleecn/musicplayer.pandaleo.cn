package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

type Error struct {
	StatusCode int    `json:"code"`
	Method     string `json:"-"`
	Path       string `json:"-"`
	Message    string `json:"message"`
	Timestamp  int64  `json:"timestamp"`
}

func newError(statusCode int, method, path, format string, args ...interface{}) Error {
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	if msg == "" {
		msg = iris.StatusText(statusCode)
	}

	return Error{
		StatusCode: statusCode,
		Method:     method,
		Path:       path,
		Message:    msg,
		Timestamp:  time.Now().Unix(),
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("[%d] %s: %s: %s", e.StatusCode, e.Method, e.Path, e.Message)
}

func (e Error) Is(target error) bool {
	if target == nil {
		return false
	}

	err, ok := target.(Error)
	if !ok {
		return false
	}

	return (err.StatusCode == e.StatusCode || e.StatusCode == 0) &&
		(err.Message == e.Message || e.Message == "")
}