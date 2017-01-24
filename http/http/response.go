package http

import (
	"encoding/json"
	"net/http"

	"vcrlibraries/domain"
	"vcrlibraries/errors"
	"vcrlibraries/loggers"
)

type Response struct {
	Data     domain.Data `json:"data"`
	HTTPCode int         `json:"-"`
	Code     string      `json:"code"`
	Message  string      `json:"message"`
}

//NewResponse 正常返回结果
func NewResponse() *Response {
	return &Response{HTTPCode: http.StatusOK, Code: "OK", Message: "OK"}
}

//NewResponseWithError 错误返回结果
func NewResponseWithError(err error) *Response {
	return &Response{
		HTTPCode: errors.CodeToHTTPCode(err),
		Code:     err.Error(),
	}
}

//NewResponseWithErrorMessage 带消息的错误返回结果
func NewResponseWithErrorMessage(err error, message string) *Response {
	response := NewResponseWithError(err)
	response.Message = message
	return response
}

//NewResponseForRedirect 带302重定向的返回结果
func NewResponseForRedirect(url string) *Response {
	return &Response{
		HTTPCode: 302,
		Message:  url,
	}
}

//Write 返回结果给用户
func (self *Response) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(self.HTTPCode)

	body, err := json.Marshal(self)
	if err != nil {
		loggers.Error.Panicln("Marshal json failed" + err.Error())
	}
	w.Write(body)
}
