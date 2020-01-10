package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//FileLogger 日志结构体
type FileLogger struct {
	logLevel    LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

//NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		logLevel:    level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}

	return f1
}

func (f *FileLogger) initFile() error {
	fullName := path.Join(f.filePath, f.fileName)
	file, err := os.OpenFile(fullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open log file failed, err:%v\n", err)
		return err
	}
	errFile, err := os.OpenFile(fullName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open error log file failed, err:%v\n", err)
		return err
	}
	//日志文件都已经打开
	f.fileObj = file
	f.errFileObj = errFile

	return nil
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] : %s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			//如果要记录的日志大于等于ERROR级别，还要在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] : %s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

//enable 告警级别开关
func (f *FileLogger) enable(loglevel LogLevel) bool {
	return f.logLevel >= loglevel
}

//Debug debug信息
func (f FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

//Info info信息
func (f FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

//Warning warning信息
func (f FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

//Error error信息
func (f FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

//Fatal fatal信息
func (f FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

//Close 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
