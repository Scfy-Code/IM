package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// APPCFG 应用的配置
var APPCFG = &AppCfg{}

// AppCfg 应用配置相关
type AppCfg struct {
	StaticDir   string            `json:"staticDir"` //静态文件目录
	ViewDir     string            `json:"viewDir"`   //模板页面目录
	DataBaseCfg `json:"dataBase"` //sql数据库的配置
	RedisCfg    `json:"redis"`    //缓存数据库的配置
	LoggerCfg   `json:"logger"`   //日志的配置
}

// DataBaseCfg 关系型数据库配置相关
type DataBaseCfg struct {
	DriverName     string `json:"driverName"`     //数据库名称
	DataSourceName string `json:"dataSourceName"` //数据库连接信息
}

// RedisCfg redis配置相关
type RedisCfg struct {
	Addrs []string `json:"addrs"` //redis连接的地址
}

// LoggerCfg 日志配置相关
type LoggerCfg struct {
	InfoLogDir string `json:"infoLogDir"` //日志存放目录
	WarnLogDir string `json:"warnLogDir"` //日志存放目录
	ErrLogDir  string `json:"errLogDir"`  //日志存放目录
}

func init() {
	data, err := ioutil.ReadFile("../cfg/application.json")
	if err != nil {
		log.Printf("读取解析配置文件错误%s", err.Error())
		os.Exit(2)
	}
	err0 := json.Unmarshal(data, APPCFG)
	if err0 != nil {
		log.Printf("配置项实例化出错%s", err0.Error())
		os.Exit(2)
	}
}
