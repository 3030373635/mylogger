package mylogger

import (
	"fmt"
	"testing"
	"time"
)

// 测试日志库
/*尚未解决问题
1. 如果使用异步日志，其他程序调用异步日志，并且其他程序过早退出，会导致异步日志无法正常写入，
因为主groutine一但挂掉，由它启动的所有子groutine也会一起结束

*/

func TestFileLog(t *testing.T) {
	log, err := NewFileLogger("info", "./", "log.log", 10*1024)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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
	log, err := NewConsoleLog("debug")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		log.Debug("这是一条Debug日志, %s写的", "qimeng")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 2)
	}
}
