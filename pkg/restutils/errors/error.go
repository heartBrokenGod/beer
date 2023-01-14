package errors

import "net/http"

type Error interface {
	error
	StatusCode() int
}

// default constructor
func New(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusInternalServerError,
	}
}
