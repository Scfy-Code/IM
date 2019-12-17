package sys

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	//只使用初始化方法
	_ "github.com/go-sql-driver/mysql"
)

var (
	sqlClients map[string]*sql.DB = make(map[string]*sql.DB)
	// RedisClient redis客户端
	RedisClient redis.UniversalClient
	// InfoLogger 常规日志对象
	InfoLogger log.Logger
	// WarnLogger 常规日志对象
	WarnLogger log.Logger
	// APP 配置对象
	APP              application = application{}
	temp             Template
	universalHandler = newautherHandler()
)

func init() {
	data, err := ioutil.ReadFile("../web/application.json")
	if err != nil {
		log.Printf("读取配置文件出错！错误信息：%s", err.Error())
		os.Exit(2)
	}
	err0 := json.Unmarshal(data, &APP)
	if err0 != nil {
		log.Printf("解析配置文件出错！错误信息：%s", err0.Error())
		os.Exit(2)
	}
	switch APP.RuntimeEnv {
	case "DEV":
		InfoLogger = *log.New(os.Stdout, "info-", log.Llongfile|log.LstdFlags)
		WarnLogger = *log.New(os.Stdout, "warn-", log.Llongfile|log.LstdFlags)
		temp = &templateDEV{
			analysisTemplateDirs(
				APP.TemplateDir,
			),
		}
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
		temp = &templatePRO{
			analysisTemplateFiles(
				analysisTemplateDirs(
					APP.TemplateDir,
				)...,
			),
		}
	}
	for index, config := range APP.SQLConfig {
		db, err0 := sql.Open(config.DriverName, config.DataSource)
		if err0 != nil {
			WarnLogger.Fatalf("创建第%d个数据源出错！错误信息：%s", index, err0.Error())
		}
		err1 := db.Ping()
		if err1 != nil {
			WarnLogger.Fatalf("连接第%d个数据源出错!错误信息：%s", index, err1.Error())
		}
		InfoLogger.Printf("创建数据源%s成功！连接信息：%s", config.ClientAlias, config.DataSource)
		sqlClients[config.ClientAlias] = db
	}
	RedisClient = redis.NewUniversalClient(APP.RedisOptions)
}
