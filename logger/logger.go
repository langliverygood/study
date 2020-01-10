package logger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//LogLevel 自定义类型
type LogLevel uint16

//定义日志级别
const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
	UNKNOWN
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLogString(logLevel LogLevel) string {
	switch logLevel {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed, err:%v\n", ok)
		return
	}
	funcName = strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
	fileName = path.Base(file)

	return funcName, fileName, lineNo
}
