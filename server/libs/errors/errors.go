package errors

import (
	"net/http"
	"strings"
)

type CustomError struct {
	Code    int
	message string
}

const (
	Unknown            = http.StatusUnprocessableEntity
	Invalid            = http.StatusBadRequest
	NotFound           = http.StatusNotFound
	AlreadyExists      = http.StatusConflict
	PermissionDenied   = http.StatusForbidden
	Unauthorized       = http.StatusUnauthorized
	Internal           = http.StatusInternalServerError
	Unavailable        = http.StatusServiceUnavailable
	NotPaid            = http.StatusPaymentRequired
	PreconditionFailed = http.StatusPreconditionFailed
	ExpectationFailed  = http.StatusExpectationFailed
)

func (e *CustomError) Error() string {
	return e.message
}

func New(c int, m string) *CustomError {
	return &CustomError{
		Code:    c,
		message: m,
	}
}

func Detect(err error) int {
	code := Unknown
	e := err.Error()
	if strings.Contains(e, "violates") {
		code = Invalid
	}
	if strings.Contains(e, "duplicate key") {
		code = AlreadyExists
	}
	if strings.Contains(e, "no rows in result") {
		code = NotFound
	}
	return code
}

func From(e error) *CustomError {
	t, ok := e.(*CustomError)
	if ok {
		return t
	}
	return New(http.StatusUnprocessableEntity, e.Error())
}

func GetCodeFrom(e error) int {
	err := From(e)
	return err.Code
}
