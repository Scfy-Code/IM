package log

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/Scfy-Code/IM/pkg"
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
