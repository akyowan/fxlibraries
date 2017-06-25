package errors

type FXError struct {
	httpCode int
	errCode  string
	msg      string
}

func NewFXError(httpCode int, code, msg string) FXError {
	return FXError{httpCode, code, msg}
}

func (e FXError) Error() string {
	return e.errCode
}

func (e FXError) HttpCode() int {
	return e.httpCode
}

func (e FXError) ErrMsg() string {
	return e.msg
}
