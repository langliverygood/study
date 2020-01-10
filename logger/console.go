package logger

import (
	"fmt"
	"os"
	"time"
)

//Logger 日志结构体
type Logger struct {
	logLevel LogLevel
}

//NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		logLevel: level,
	}
}

func (l Logger) enable(loglevel LogLevel) bool {
	return l.logLevel >= loglevel
}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Fprintf(os.Stdout, "[%s] [%s] [%s:%s:%d] : %s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}

//Debug debug信息
func (l Logger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		log(DEBUG, format, a...)
	}
}

//Info info信息
func (l Logger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format, a...)
	}
}

//Warning warning信息
func (l Logger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format, a...)
	}
}

//Error error信息
func (l Logger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, a...)
	}
}

//Fatal fatal信息
func (l Logger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL, format, a...)
	}
}
