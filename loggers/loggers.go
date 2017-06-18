package loggers

import (
	"bytes"
	"io"
	"log"
	"os"
)

var (
	Error *log.Logger
	Info  *log.Logger
	Debug *log.Logger
	Warn  *log.Logger
)

func init() {
	Error = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	Info = log.New(os.Stderr, "[Info] ", log.LstdFlags|log.Lshortfile)
	Warn = log.New(os.Stderr, "[Warn] ", log.LstdFlags|log.Lshortfile)
	Debug = log.New(os.Stderr, "[Debug] ", log.LstdFlags|log.Lshortfile)
}

func SetOutput(w io.Writer) {
	Error.SetOutput(w)
	Info.SetOutput(w)
	Debug.SetOutput(w)
	Warn.SetOutput(w)
}

func CaptureLog(f func()) string {
	var buf bytes.Buffer
	SetOutput(&buf)
	f()
	SetOutput(os.Stderr)
	return buf.String()
}
