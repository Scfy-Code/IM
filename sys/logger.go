package sys

import (
	"io"
	"log"
	"os"
	"time"
)

//
func init() {
	switch APP.RuntimeEnv {
	case "DEV":
		InfoLogger = *log.New(os.Stdout, "info-", log.Llongfile|log.LstdFlags)
		WarnLogger = *log.New(os.Stdout, "warn-", log.Llongfile|log.LstdFlags)
	case "PRO":
		infoLog, err0 := os.OpenFile(APP.LoggerDir+"/info-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err0 != nil {
			log.Fatalf("设置常规日志文件出错!错误信息：%s", err0.Error())
		}
		warnLog, err1 := os.OpenFile(APP.LoggerDir+"/warn-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err1 != nil {
			log.Fatalf("设置错误日志文件出错!错误信息：%s", err1.Error())
		}
		InfoLogger = *log.New(io.MultiWriter(os.Stdout, infoLog), "info-", log.Llongfile|log.LstdFlags)
		WarnLogger = *log.New(io.MultiWriter(os.Stdout, warnLog), "warn-", log.Llongfile|log.LstdFlags)
	}
}
