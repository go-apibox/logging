package logging

import (
	"os"

	"github.com/op/go-logging"
)

type Logger struct {
	*logging.Logger
}

var infoBackendLeveled logging.LeveledBackend
var errorBackendLeveled logging.LeveledBackend

func init() {
	accessOut := os.Stdout
	accessLogFile := os.Getenv("ACCESS_LOG")
	if accessLogFile != "" {
		f, err := os.OpenFile(accessLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			accessOut = f
		} else {
			// 忽略输出
			f, err = os.OpenFile(os.DevNull, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err == nil {
				accessOut = f
			}
		}
	}
	infoFormatter := logging.MustStringFormatter(
		"%{color}%{time:01-02 15:04:05.000} %{module} » %{level:.4s} %{id:03x}%{color:reset} %{message}",
	)
	infoBackend := logging.NewLogBackend(accessOut, "", 0)
	infoBackendFormatter := logging.NewBackendFormatter(infoBackend, infoFormatter)
	infoBackendLeveled = logging.AddModuleLevel(infoBackendFormatter)

	errOut := os.Stderr
	errorLogFile := os.Getenv("ERROR_LOG")
	if errorLogFile != "" {
		f, err := os.OpenFile(errorLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			errOut = f
		} else {
			// 忽略输出
			f, err = os.OpenFile(os.DevNull, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err == nil {
				errOut = f
			}
		}
	}
	errorFormatter := logging.MustStringFormatter(
		"%{color}%{time:01-02 15:04:05.000} %{module} » %{level:.4s} %{id:03x}%{color:reset} [%{shortfile} %{shortfunc}] %{message}",
	)
	errorBackend := logging.NewLogBackend(errOut, "", 0)
	errorBackendFormatter := logging.NewBackendFormatter(errorBackend, errorFormatter)
	errorBackendLeveled = logging.AddModuleLevel(errorBackendFormatter)

	logging.SetBackend(infoBackendLeveled, errorBackendLeveled)
}

// NewLogger returns a new Logger instance
func NewLogger(module string) *Logger {
	infoBackendLeveled.SetLevel(logging.DEBUG, module)
	errorBackendLeveled.SetLevel(logging.WARNING, module)
	return &Logger{logging.MustGetLogger(module)}
}
