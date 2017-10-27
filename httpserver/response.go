package httpserver

import (
	"encoding/json"
	"encoding/xml"
	"fxlibraries/errors"
	"fxlibraries/loggers"
	"net/http"
	"strings"
)

type Response struct {
	HTTPCode int         `json:"-"`
	Code     string      `json:"errno"`
	Msg      string      `json:"message"`
	Header   http.Header `json:"-"`
	Data     interface{} `json:"data,omitempty"`
	IsWx     bool        `json:"-"`
}

type XMLResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

// NewResponse
func NewResponse() *Response {
	resp := &Response{
		HTTPCode: http.StatusOK,
		Code:     "OK",
		Msg:      "Success",
		Header:   make(http.Header),
	}
	resp.Header.Set("Content-Type", "application/json")
	return resp
}

// NewResponseWithError
func NewResponseWithError(err errors.FXError) *Response {
	resp := &Response{
		HTTPCode: err.HttpCode(),
		Code:     err.Error(),
		Msg:      err.ErrMsg(),
		Header:   make(http.Header),
	}
	resp.Header.Set("Content-Type", "application/json")
	return resp
}

// NewResponseForRedirect
func NewResponseForRedirect(url string) *Response {
	resp := &Response{
		HTTPCode: http.StatusFound,
		Header:   make(http.Header),
	}
	resp.Header.Set("Location", url)

	return resp
}

// Write
func (r *Response) Write(w http.ResponseWriter) {
	if r.IsWx {
		resp := &XMLResponse{
			ReturnCode: "SUCCESS",
			ReturnMsg:  "OK",
		}
		body, _ := xml.Marshal(resp)
		dataStr := strings.Replace(string(body), "XMLResponse", "xml", -1)
		w.Write([]byte(dataStr))
		return

	}
	for k, vv := range r.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(r.HTTPCode)
	body, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}
