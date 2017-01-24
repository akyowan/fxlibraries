package loggers

import (
	"encoding/json"
	log "github.com/cihub/seelog"
	"rdslibraries/http"
)

var runLogger log.LoggerInterface
var accessLogger log.LoggerInterface

func init() {
	defaultConfig := `
	<seelog>
		<output>
			<console />
		</output>
	</seelog>
	`
	var err error
	runLogger, err = log.LoggerFromConfigAsBytes([]byte(defaultConfig))
	if err != nil {
		panic(err)
	}

	accessConfig := `
	<seelog>
		<output>
		</output>
	</seelog>
	`
	accessLogger, err = log.LoggerFromConfigAsBytes([]byte(accessConfig))
	if err != nil {
		panic(err)
	}
}

func Tracef(format string, params ...interface{}) {
	return runLogger.Tracef(format, params...)
}

func Debugf(format string, params ...interface{}) {
	return runLogger.Debugf(format, params...)
}

func Infof(format string, params ...interface{}) {
	return runLogger.Infof(format, params...)
}

func Warnf(format string, params ...interface{}) error {
	return runLogger.Warnf(format, params...)
}

func Errorf(format string, params ...interface{}) error {
	return runLogger.Errorf(format, params...)
}

func Criticalf(format string, params ...interface{}) error {
	return runLogger.Criticalf(format, params...)
}

func Trace(v ...interface{}) {
	return runLogger.Trace(v...)
}

func Debug(v ...interface{}) {
	return runLogger.Debug(v...)
}

func Info(v ...interface{}) {
	return runLogger.Info(v...)
}

func Warn(v ...interface{}) error {
	return runLogger.Warn(v...)
}

func Error(v ...interface{}) error {
	return runLogger.Error(v...)
}

func Critical(v ...interface{}) error {
	return runLogger.Critical(v...)
}

// Close flushes all the messages in the logger and closes it. It cannot be used after this operation.
func Close() {
	runLogger.Close()
	accessLog.Close()
}

// Flush flushes all the messages in the logger.
func Flush() {
	runLogger.Flush()
	accessLog.Flush()
}

// Run return the runLogger
func Run() log.LoggerInterface {
	return runLogger
}

func Access(req *http.Request, resp *http.Response, handTime int64) {
	accessLog := map[string]interface{}{
		"handTime": handTime,
		"request":  req,
		"response": response,
	}
	accessStr, err := json.Marshal(accessLog)
	if err != nil {
		loggers.Errorf("Marshal access log error:%s", err.Error())
	}
}
