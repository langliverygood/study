package logger

import (
	"fmt"
	"os"
	"time"
)

//ConsoleLogger 日志结构体
type ConsoleLogger struct {
	logLevel LogLevel
}

//NewConsoleLogger 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		logLevel: level,
	}
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(os.Stdout, "[%s] [%s] [%s:%s:%d] : %s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

//enable 告警级别开关
func (c ConsoleLogger) enable(loglevel LogLevel) bool {
	return c.logLevel >= loglevel
}

//Debug debug信息
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

//Info info信息
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

//Warning warning信息
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

//Error error信息
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

//Fatal fatal信息
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
