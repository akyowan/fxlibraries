package httpserver

import (
	"encoding/json"
	"fxlibraries/errors"
	"net/http"
)

type Response struct {
	HTTPCode int         `json:"-"`
	Code     string      `json:"errno"`
	Msg      string      `json:"message"`
	Header   http.Header `json:"-"`
	Data     interface{} `json:"data,omitempty"`
}

// NewResponse
func NewResponse() *Response {
	return &Response{
		HTTPCode: http.StatusOK,
		Code:     "OK",
		Msg:      "Success",
	}
}

// NewResponseWithError
func NewResponseWithError(err errors.FXError) *Response {
	return &Response{
		HTTPCode: err.HttpCode(),
		Code:     err.Error(),
		Msg:      err.ErrMsg(),
	}
}

// NewResponseForRedirect
func NewResponseForRedirect(url string) *Response {
	resp := &Response{
		HTTPCode: http.StatusFound,
	}
	resp.Header.Set("Location", url)

	return resp
}

// Write
func (r *Response) Write(w http.ResponseWriter) {
	for k, vv := range r.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.HTTPCode)
	body, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}
