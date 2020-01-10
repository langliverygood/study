package main

import (
	"time"

	"../logger"
)

func main() {
	log := logger.NewFileLogger("FAtal", "./", "logger_test", 10*1024)
	for {
		log.Debug("这是一条测试日志")
		log.Info("这是一条测试日志")
		log.Warning("这是一条测试日志")
		log.Error("这是一条测试日志")
		log.Fatal("这是一条测试日志")
		time.Sleep(time.Second)
	}
}
