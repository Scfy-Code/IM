package log

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	infoOut, warnOut io.Writer
)

// NewInfoLogger 创建一个info级别的日志
func NewInfoLogger() *log.Logger {
	return log.New(io.MultiWriter(os.Stdout, infoOut), "info-", log.Llongfile|log.LstdFlags)
}

// NewWarnLogger 创建一个警告级别的日志对象
func NewWarnLogger() *log.Logger {
	return log.New(io.MultiWriter(os.Stdout, warnOut), "warn-", log.Llongfile|log.LstdFlags)
}

// RegistLogDir 注册日志l
func RegistLogDir(logDir string) {
	infoLog, err0 := os.OpenFile(logDir+"/info-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	warnLog, err1 := os.OpenFile(logDir+"/warn-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err0 != nil || err1 != nil {
		os.Exit(2)
	}
	infoOut = infoLog
	warnOut = warnLog
}
