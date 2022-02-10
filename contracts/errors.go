package contracts

import (
	"net/http"
)

type errorDetails struct {
	message string
	status  int
}

var errorObjects = map[string]errorDetails{
	ErrorBadRequest: {
		message: "bad request",
		status:  http.StatusBadRequest,
	},
	ErrorUnauthorized: {
		message: "unauthorized",
		status:  http.StatusUnauthorized,
	},
	ErrorForbidden: {
		message: "forbidden",
		status:  http.StatusForbidden,
	},
	ErrorNotFound: {
		message: "not found",
		status:  http.StatusNotFound,
	},
	ErrorMethodNotAllowed: {
		message: "not allowed",
		status:  http.StatusMethodNotAllowed,
	},
	ErrorNotAcceptable: {
		message: "not acceptable",
		status:  http.StatusNotAcceptable,
	},
	ErrorRequestTimeout: {
		message: "request timeout",
		status:  http.StatusRequestTimeout,
	},
	ErrorConflict: {
		message: "conflict",
		status:  http.StatusConflict,
	},
	ErrorUnprocessableEntity: {
		message: "unprocessable entity",
		status:  http.StatusUnprocessableEntity,
	},
	ErrorUpgradeRequired: {
		message: "upgrade required",
		status:  http.StatusUpgradeRequired,
	},
	ErrorTooManyRequests: {
		message: "too many requests",
		status:  http.StatusTooManyRequests,
	},
	ErrorInternalServer: {
		message: "server error",
		status:  http.StatusInternalServerError,
	},
	ErrorServiceUnavailable: {
		message: "service unavailable",
		status:  http.StatusServiceUnavailable,
	},
}
