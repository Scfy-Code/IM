package log

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

// InfoLog 用于普通的日志输出
func InfoLog(v ...interface{}) {
	info.Println(v...)
}

// WarnLog 用于异常的日志输出
func WarnLog(v ...interface{}) {
	info.Println(v...)
}

// ErrLog 用于错误的日志输出
func ErrLog(v ...interface{}) {
	info.Println(v...)
}
func init() {
	infoFile, err0 := os.OpenFile(config.APPCFG.LoggerCfg.InfoLogDir+time.Now().Format("2006-01-02")+"info.log", os.O_APPEND|os.O_CREATE, os.ModeAppend)
	warnFile, err1 := os.OpenFile(config.APPCFG.LoggerCfg.WarnLogDir+time.Now().Format("2006-01-02")+"warn.log", os.O_APPEND|os.O_CREATE, os.ModeAppend)
	errFile, err2 := os.OpenFile(config.APPCFG.LoggerCfg.ErrLogDir+time.Now().Format("2006-01-02")+"error.log", os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err0 != nil || err1 != nil || err2 != nil {
		log.Println("读取日志文件失败")
		os.Exit(2)
	}
	info = log.New(io.MultiWriter(os.Stdout, infoFile), "info-", log.Ldate|log.Ltime|log.Llongfile)
	warn = log.New(io.MultiWriter(os.Stdout, warnFile), "warn-", log.Ldate|log.Ltime|log.Llongfile)
	err = log.New(io.MultiWriter(os.Stdout, errFile), "error-", log.Ldate|log.Ltime|log.Llongfile)
}
