package mylogger

import (
	"testing"
	"time"
)

// 测试日志库

func TestFileLog(t *testing.T) {
	log := NewFileLogger("info", "./", "log.log", 10*1024)
	for {
		log.Debug("这是一条Debug日志, %s写的", "qimeng")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 2)
	}
}

func TestConsoleLog(t *testing.T) {
	log := NewConsoleLog("debug")
	for {
		log.Debug("这是一条Debug日志, %s写的", "qimeng")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 2)
	}
}
