package main

import (
	"study/logger"
	"time"
)

func main() {
	logger := logger.NewLog("eRrOR")
	for {
		logger.Debug("这是一条测试日志！%d", 123)
		logger.Info("这是一条测试日志！")
		logger.Warning("这是一条测试日志！")
		logger.Error("这是一条测试日志！")
		logger.Fatal("这是一条测试日志！")
		time.Sleep(time.Second)
	}
}
