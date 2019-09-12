package app

import "log"

var (
	// InfoLogger 常规日志
	InfoLogger *log.Logger
	// WarnLogger 警告日志
	WarnLogger *log.Logger
	// ErrorLogger 错误日志
	ErrorLogger *log.Logger
)
