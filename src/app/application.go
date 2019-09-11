package app

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	cache "github.com/go-redis/redis"
	// 只是用init()方法注册驱动
	_ "github.com/go-sql-driver/mysql"
)

// StaticDir 静态资源目录
var StaticDir string

//应用配置
type application struct {
	StaticDir   string     `json:"staticDir"`   //静态资源目录
	TemplateDir string     `json:"templateDir"` //模板目录
	UploadDir   string     `json:"uploadDir"`   //上传文件地址
	DataSource  dataSource `json:"dataSource"`  //数据源配置
	Logger      logger     `json:"logger"`      //日志配置
}

//数据源配置
type dataSource struct {
	Mysql mysql `json:"mysql"` //mysql数据库配置
	Redis redis `json:"redis"` //redis数据库配置
}

func newApplication(configFile string) *application {
	//读取配置文件
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("读取解析配置文件错误%s", err.Error())
		os.Exit(2)
	}
	//将配置文件的序列化成对象
	app := &application{}
	err0 := json.Unmarshal(data, app)
	if err0 != nil {
		log.Printf("配置项实例化出错%s", err0.Error())
		os.Exit(2)
	}
	return app
}
func (a application) newMysqlClient() *sql.DB {
	driveName := a.DataSource.Mysql.DriverName
	dataSourceName := a.DataSource.Mysql.UserName + ":" + a.DataSource.Mysql.Password + "@tcp(" + a.DataSource.Mysql.Host + ")/" + a.DataSource.Mysql.DataBase + "?charset=" + a.DataSource.Mysql.Charset
	db, err := sql.Open(driveName, dataSourceName)
	if err != nil {
		ErrorLogger.Printf("连接数据库出错！错误信息:%s", err.Error())
		os.Exit(2)
	}
	return db
}
func (a application) newRedisClient() cache.UniversalClient {
	if a.DataSource.Redis.Cluster {
		return cache.NewClusterClient(a.DataSource.Redis.ClusterOptions)
	}
	return cache.NewClient(a.DataSource.Redis.StandOptions)
}

func (a application) newLoggerClient() (*log.Logger, *log.Logger, *log.Logger) {
	infoLogFile, err0 := os.OpenFile(a.Logger.LoggerDir+"info/info-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	warnLogFile, err1 := os.OpenFile(a.Logger.LoggerDir+"warn/warn-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	errorLogFile, err2 := os.OpenFile(a.Logger.LoggerDir+"error/error-"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err0 != nil || err1 != nil || err2 != nil {
		log.Printf("创建日志文件失败！错误信息：%s;%s;%s;", err0.Error(), err1.Error(), err2.Error())
	}
	return log.New(io.MultiWriter(os.Stdout, infoLogFile), "info-", log.LstdFlags|log.Llongfile), log.New(io.MultiWriter(os.Stdout, warnLogFile), "warn-", log.LstdFlags|log.Llongfile), log.New(io.MultiWriter(os.Stderr, errorLogFile), "error-", log.LstdFlags|log.Llongfile)
}

func init() {
	appc := newApplication("E:\\SCFY\\文档\\VsCode\\scfy-im\\cfg\\application.json")
	InfoLogger, WarnLogger, ErrorLogger = appc.newLoggerClient()
	MysqlClient = appc.newMysqlClient()
	RedisClient = appc.newRedisClient()
	newTemplateMap(appc.TemplateDir)
	StaticDir = appc.StaticDir
}
