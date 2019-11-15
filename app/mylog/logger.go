package mylog

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/Scfy-Code/IM/app"
)

var (
	infoOut, warnOut io.Writer
)

// NewInfoLogger 创建一个info级别的日志
func NewInfoLogger() *log.Logger {
	var logger = log.New(os.Stdout, "info-", log.Llongfile|log.LstdFlags)
	switch pkg.APP.RuntimeEnv {
	case "DEV":
		return logger
	case "PRO":
		logger.SetOutput(io.MultiWriter(os.Stdout, infoOut))
		return logger
	default:
		logger.SetOutput(io.MultiWriter(os.Stdout, infoOut))
		return logger
	}
}

// NewWarnLogger 创建一个警告级别的日志对象
func NewWarnLogger() *log.Logger {
	var logger = log.New(os.Stdout, "warn-", log.Llongfile|log.LstdFlags)
	switch pkg.APP.RuntimeEnv {
	case "DEV":
		return logger
	case "PRO":
		logger.SetOutput(io.MultiWriter(os.Stdout, warnOut))
		return logger
	default:
		logger.SetOutput(io.MultiWriter(os.Stdout, warnOut))
		return logger
	}
}

//
func init() {
	infoLog, err0 := os.OpenFile(pkg.APP.LoggerDir+"/info-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err0 != nil {
		log.Fatalf("设置常规日志文件出错!错误信息：%s", err0.Error())
	}
	warnLog, err1 := os.OpenFile(pkg.APP.LoggerDir+"/warn-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err1 != nil {
		log.Fatalf("设置错误日志文件出错!错误信息：%s", err0.Error())
	}
	infoOut = infoLog
	warnOut = warnLog
}
