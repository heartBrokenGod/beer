package errors

type ErrorImpl struct {
	msg        string
	statusCode int
}

func (e *ErrorImpl) StatusCode() int {
	return e.statusCode
}

func (e *ErrorImpl) Error() string {
	return e.msg
}
