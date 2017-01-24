package errors

import (
	"net/http"
)

var (
	BadRequest          = NewRDSError(http.StatusBadRequest, "BAD_REQUEST", http.StatusText(http.StatusBadRequest))
	Unauthorized        = NewRDSError(http.StatusUnauthorized, "UNAUTHORIZED", http.StatusText(http.StatusUnauthorized))
	Forbidden           = NewRDSError(http.StatusForbidden, "FORBIDDEN", http.StatusText(http.StatusForbidden))
	InternalServerError = NewRDSError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", http.StatusText(http.StatusInternalServerError))
)
