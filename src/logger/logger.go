package logger

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/Scfy-Code/scfy-im/config"
)

var (
	info *log.Logger //用于正常的日志输出
	warn *log.Logger //用于异常的日志输出
	err  *log.Logger //用于错误的日志输出
)

// InfoPrintf 普通日志格式化输出
func InfoPrintf(format string, v ...interface{}) {
	info.Printf(format, v...)
}

// InfoPrintln 普通日志换行输出
func InfoPrintln(v ...interface{}) {
	info.Println(v...)
}

// WarnPrintf 用于异常的日志格式化输出
func WarnPrintf(format string, v ...interface{}) {
	warn.Printf(format, v...)
}

// WarnPrintln 用于异常的日志换行输出
func WarnPrintln(v ...interface{}) {
	warn.Println(v...)
}

// ErrPrintf 用于错误的日志格式化输出
func ErrPrintf(format string, v ...interface{}) {
	err.Printf(format, v...)
}

// ErrPrintln 用于错误的换行输出
func ErrPrintln(v ...interface{}) {
	err.Println(v...)
}
func init() {
	infoFile, err0 := os.OpenFile(config.APPCFG.LoggerCfg.InfoLogDir+"info-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	warnFile, err1 := os.OpenFile(config.APPCFG.LoggerCfg.WarnLogDir+"warn-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	errFile, err2 := os.OpenFile(config.APPCFG.LoggerCfg.ErrLogDir+"err-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err0 != nil || err1 != nil || err2 != nil {
		log.Printf("读取日志文件失败!错误信息：%s;%s;%s", err0.Error(), err1.Error(), err2.Error())
		os.Exit(2)
	}
	info = log.New(io.MultiWriter(os.Stdout, infoFile), "info-", log.Ldate|log.Ltime|log.Llongfile)
	warn = log.New(io.MultiWriter(os.Stdout, warnFile), "warn-", log.Ldate|log.Ltime|log.Llongfile)
	err = log.New(io.MultiWriter(os.Stdout, errFile), "err-", log.Ldate|log.Ltime|log.Llongfile)
}
