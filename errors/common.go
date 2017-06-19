package errors

import (
	"net/http"
)

var (
	BadRequest          = NewFXError(http.StatusBadRequest, "BAD_REQUEST", http.StatusText(http.StatusBadRequest))
	Unauthorized        = NewFXError(http.StatusUnauthorized, "UNAUTHORIZED", http.StatusText(http.StatusUnauthorized))
	Forbidden           = NewFXError(http.StatusForbidden, "FORBIDDEN", http.StatusText(http.StatusForbidden))
	InternalServerError = NewFXError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", http.StatusText(http.StatusInternalServerError))
)
