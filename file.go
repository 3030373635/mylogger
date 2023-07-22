package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	maxFileSize int
	logChan     chan *LogMsg // 日志通道，异步写日志
}

type LogMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

func NewFileLogger(levelStr, filePath, fileName string, maxFileSize int) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       level,
		fileName:    fileName,
		filePath:    filePath,
		maxFileSize: maxFileSize,
		logChan:     make(chan *LogMsg, 50000),
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	// 开启5个协程从管道读取消息写日志(有问题。某个协程关闭了文件。其他协程正在读文件就会有问题)
	// for i:=0;i<3;i++{
	// 	go f1.writeLogBackGround()
	// }
	go f1.writeLogBackGround()

	return f1
}

func (f *FileLogger) checkFile(fileObj *os.File) bool {
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Println("获取文件状态失败")
		return false
	}
	fileSize := fileInfo.Size()
	return fileSize > int64(f.maxFileSize)
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("初始化日志文件失败")
		return err
	}
	f.fileObj = fileObj
	return err

}

func (f *FileLogger) writeLogBackGround() {
	for {
		if f.checkFile(f.fileObj) {
			// 切分日志文件
			// 1. 关闭当前日志文件
			// 2. 重命名原文件 ***.log.back20201020153030
			// 3. 打开新的文件
			f.fileObj.Close()

			newFileName := fmt.Sprintf("%s.back%s", f.fileName, time.Now().Format("20060102150405"))
			oldPath := path.Join(f.filePath, f.fileName)
			newPath := path.Join(f.filePath, newFileName)
			os.Rename(oldPath, newPath)

			fileObj, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("备份文件失败!")
				return
			}
			f.fileObj = fileObj
		}
		select {
		case logTmp := <-f.logChan:
			fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, getLogLevel(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
		default:
		}
	}

}

// 写日志
func (f *FileLogger) log(level LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now().Format("2006-01-02 15:04:05")
	fileName, funcName, lineNo := getCallInfo(3)
	logTmp := &LogMsg{
		level:     level,
		msg:       msg,
		funcName:  funcName,
		fileName:  fileName,
		timestamp: now,
		line:      lineNo,
	}
	//保证程序顺畅执行，不阻塞
	select {
	case f.logChan <- logTmp:
	default:
	}

}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	if f.enable(DEBUG) {
		f.log(DEBUG, format, a...)
	}

}
func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		f.log(INFO, format, a...)
	}
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	if f.enable(WARNING) {
		f.log(WARNING, format, a...)
	}

}

func (f *FileLogger) Error(format string, a ...interface{}) {
	if f.enable(ERROR) {
		f.log(ERROR, format, a...)
	}

}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	if f.enable(FATAL) {
		f.log(FATAL, format, a...)
	}

}
