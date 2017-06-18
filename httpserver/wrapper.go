package httpserver

import (
	"fxlibraries/errors"
	"fxlibraries/loggers"
	"net/http"
	"net/url"
	"time"
)

// handlerWrapper Every http api hundler will be processed by this wrapper
func HandlerWrapper(f HandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				loggers.Error.Println(r)
				NewResponseWithError(errors.InternalServerError)
			}

		}()
		startTime := time.Now()

		request := NewRequest(r)
		response := f(request)
		response.Write(w)

		loggers.Info.Printf(`%3d %s %s %s %v`,
			response.HTTPCode,
			response.Code,
			request.Method,
			getPathAndQuery(r.URL),
			time.Since(startTime),
		)
	}
}

// getPathAndQuery Combile path and raw query
func getPathAndQuery(u *url.URL) string {
	if u.RawQuery == "" {
		return u.Path
	} else {
		rawQuery, _ := url.QueryUnescape(u.RawQuery)
		return u.Path + "?" + rawQuery
	}
}
