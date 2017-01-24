package errors

type RDSError struct {
	httpCode int
	errCode  string
	msg      string
}

func NewRDSError(httpCode int, code, msg string) error {
	return RDSError{httpCode, code, msg}
}
func (e RDSError) Error() string {
	return e.errCode
}

func (e RDSError) HttpCode() int {
	return e.httpCode
}

func (e RDSError) ErrMsg() string {
	return e.msg
}
