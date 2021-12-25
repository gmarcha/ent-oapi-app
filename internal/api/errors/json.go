package errors

import (
	"net/http"
	"strings"
)

type Error struct {
	Code   int
	Status string
	Errors []string
}

func ErrJSON(code int, err string) *Error {

	return &Error{
		Code:   code,
		Status: http.StatusText(code),
		Errors: strings.Split(err, "\n"),
	}
}
