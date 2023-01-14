package errors

import "net/http"

// errors for 400 status codes
func NewErrorBadRequest(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusContinue,
	}
}

func NewErrorUnauthorized(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusSwitchingProtocols,
	}
}
func NewErrorPaymentRequired(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusPaymentRequired,
	}
}

func NewErrorForbidden(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusForbidden,
	}
}

func NewErrorNotFound(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusNotFound,
	}
}

func NewErrorMethodNotAllowed(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusMethodNotAllowed,
	}
}
func NewErrorNotAcceptable(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusNotAcceptable,
	}
}

func NewErrorProxyAuthRequired(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusProxyAuthRequired,
	}
}

func NewErrorRequestTimeout(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusRequestTimeout,
	}
}

func NewErrorConflict(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusConflict,
	}
}
func NewErrorGone(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusGone,
	}
}

func NewErrorLengthRequired(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusLengthRequired,
	}
}
func NewErrorPreconditionFailed(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusPreconditionFailed,
	}
}

func NewErrorRequestEntityTooLarge(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusRequestEntityTooLarge,
	}
}
func NewErrorRequestURITooLong(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusRequestURITooLong,
	}
}

func NewErrorUnsupportedMediaType(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusUnsupportedMediaType,
	}
}

func NewErrorRequestedRangeNotSatisfiable(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusRequestedRangeNotSatisfiable,
	}
}

func NewErrorExpectationFailed(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusExpectationFailed,
	}
}
func NewErrorTeapot(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusTeapot,
	}
}

func NewErrorMisdirectedRequest(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusMisdirectedRequest,
	}
}
func NewErrorUnprocessableEntity(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusUnprocessableEntity,
	}
}

func NewErrorLocked(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusLocked,
	}
}
func NewErrorFailedDependency(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusFailedDependency,
	}
}

func NewErrorTooEarly(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusTooEarly,
	}
}

func NewErrorUpgradeRequired(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusUpgradeRequired,
	}
}

func NewErrorPreconditionRequired(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusPreconditionRequired,
	}
}
func NewErrorTooManyRequests(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusTooManyRequests,
	}
}

func NewErrorRequestHeaderFieldsTooLarge(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusRequestHeaderFieldsTooLarge,
	}
}
func NewErrorUnavailableForLegalReasons(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusUnavailableForLegalReasons,
	}
}

// errors for 500 status codes
func NewErrorInternalServerError(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusInternalServerError,
	}
}
func NewErrorNotImplemented(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusNotImplemented,
	}
}

func NewErrorBadGateway(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusBadGateway,
	}
}

func NewErrorServiceUnavailable(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusServiceUnavailable,
	}
}

func NewErrorGatewayTimeout(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusGatewayTimeout,
	}
}
func NewErrorHTTPVersionNotSupported(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusHTTPVersionNotSupported,
	}
}

func NewErrorVariantAlsoNegotiates(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusVariantAlsoNegotiates,
	}
}
func NewErrorInsufficientStorage(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusInsufficientStorage,
	}
}

func NewErrorLoopDetected(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusLoopDetected,
	}
}
func NewErrorNotExtended(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusNotExtended,
	}
}

func NewErrorNetworkAuthenticationRequired(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusNetworkAuthenticationRequired,
	}
}
