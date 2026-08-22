package response

import "net/http"

type Exception struct {
	StatusCode int
	Message    string
}

func (e *Exception) Error() string { return e.Message }

func NewBadRequestException(msg ...string) *Exception {
	return newException(http.StatusBadRequest, msg...)
}

func NewUnauthorizedException(msg ...string) *Exception {
	return newException(http.StatusUnauthorized, msg...)
}

func NewForbiddenException(msg ...string) *Exception {
	return newException(http.StatusForbidden, msg...)
}

func NewNotFoundException(msg ...string) *Exception {
	return newException(http.StatusNotFound, msg...)
}

func NewInternalServerErrorException(msg ...string) *Exception {
	return newException(http.StatusInternalServerError, msg...)
}

func newException(code int, msg ...string) *Exception {
	m := http.StatusText(code)
	if len(msg) > 0 && msg[0] != "" {
		m = msg[0]
	}
	return &Exception{StatusCode: code, Message: m}
}
