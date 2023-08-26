// 往终端写
package mylogger

import (
	"fmt"
	"time"
)

// 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// 构造函数
func NewConsoleLog(levelStr string) (ConsoleLogger, error) {

	level, err := parseLogLevel(levelStr)
	cl := ConsoleLogger{}
	if err != nil {
		return cl, err
	} else {
		cl.Level = level
		return cl, nil
	}
}

func (c ConsoleLogger) log(level LogLevel, format string, a ...interface{}) {
	if c.enable(level) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		fileName, funcName, lineNo := getCallInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogLevel(level), fileName, funcName, lineNo, msg)
	}

}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)

}
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)

}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)

}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)

}
