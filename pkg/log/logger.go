package log

import (
	"log"
	"os"
)

// NewInfoLogger 创建一个info级别的日志
func NewInfoLogger() *log.Logger {
	return log.New(os.Stdout, "info-", log.Llongfile|log.LstdFlags)
}

// NewWarnLogger 创建一个警告级别的日志对象
func NewWarnLogger() *log.Logger {
	return log.New(os.Stdout, "warn-", log.Llongfile|log.LstdFlags)
}
