package sys

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
	return log.New(infoOut, "info-", log.Llongfile|log.LstdFlags)
}

// NewWarnLogger 创建一个警告级别的日志对象
func NewWarnLogger() *log.Logger {
	return log.New(warnOut, "warn-", log.Llongfile|log.LstdFlags)
}

//
func init() {
	switch APP.RuntimeEnv {
	case "DEV":
		infoOut = os.Stdout
		warnOut = os.Stdout
	case "PRO":
		infoLog, err0 := os.OpenFile(APP.LoggerDir+"/info-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err0 != nil {
			log.Fatalf("设置常规日志文件出错!错误信息：%s", err0.Error())
		}
		warnLog, err1 := os.OpenFile(APP.LoggerDir+"/warn-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err1 != nil {
			log.Fatalf("设置错误日志文件出错!错误信息：%s", err1.Error())
		}
		infoOut = io.MultiWriter(os.Stdout, infoLog)
		warnOut = io.MultiWriter(os.Stdout, warnLog)
	default:
		infoLog, err0 := os.OpenFile(APP.LoggerDir+"/info-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err0 != nil {
			log.Fatalf("设置常规日志文件出错!错误信息：%s", err0.Error())
		}
		warnLog, err1 := os.OpenFile(APP.LoggerDir+"/warn-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err1 != nil {
			log.Fatalf("设置错误日志文件出错!错误信息：%s", err1.Error())
		}
		infoOut = io.MultiWriter(os.Stdout, infoLog)
		warnOut = io.MultiWriter(os.Stdout, warnLog)
	}
}