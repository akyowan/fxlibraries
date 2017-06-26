package errors

import (
	"net/http"
)

var (
	BadRequest          = NewFXError(http.StatusBadRequest, "BAD_REQUEST", http.StatusText(http.StatusBadRequest))
	Unauthorized        = NewFXError(http.StatusUnauthorized, "UNAUTHORIZED", http.StatusText(http.StatusUnauthorized))
	Forbidden           = NewFXError(http.StatusForbidden, "FORBIDDEN", http.StatusText(http.StatusForbidden))
	InternalServerError = NewFXError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", http.StatusText(http.StatusInternalServerError))
	NotFound            = NewFXError(http.StatusNotFound, "NOT_FOUND", http.StatusText(http.StatusNotFound))
)

func NewBadRequest(err string) FXError {
	return NewFXError(http.StatusBadRequest, "BAD_REQUEST", err)
}

func NewNotFound(err string) FXError {
	return NewFXError(http.StatusNotFound, "NOT_FOUND", err)
}

var (
	ParameterError = NewBadRequest("Invalid parameter")
)
