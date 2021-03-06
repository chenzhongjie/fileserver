package log

import (
	"fileserver/version"
	"github.com/chenzhongjie/fileLogger"
)

var Log Logger

type Logger interface {
	Debug(format string, params ...interface{})
	Info(format string, params ...interface{})
	Warn(format string, params ...interface{})
	Error(format string, params ...interface{})
}

func init() {
	logPath := "logs"
	flog := fileLogger.NewDefaultLogger(logPath, version.NAME+".log", true)
	flog.SetLogLevel(fileLogger.INFO)
	Log = flog
}