// 自定义日志库
package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

const (
	Invaild LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func getCallInfo(skip int) (fileName, funcName string, lineNo int) {
	// 获取调用日志调用信息，包括行数、文件名、函数等
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("获取行数失败")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return

}

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
	default:
		err := errors.New("无效的日志级别")
		return Invaild, err
	}
}
func getLogLevel(level LogLevel) string {
	switch level {
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
