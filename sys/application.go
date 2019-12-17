package sys

import (
	"database/sql"

	"github.com/go-redis/redis"
)

type application struct {
	Port         string                  `json:"port"`         //监听端口
	RuntimeEnv   string                  `json:"runtimeEnv"`   //运行环境
	TemplateDir  string                  `json:"templateDir"`  //视图目录
	StaticDir    string                  `json:"staticDir"`    //静态文件目录
	UploadDir    string                  `json:"uploadDir"`    //上传文件目录
	LoggerDir    string                  `json:"loggerDir"`    //日志目录
	SQLConfig    []sqlConfig             `json:"dataSources"`  //数据源配置信息
	RedisOptions *redis.UniversalOptions `json:"redisOptions"` //缓存配置
}

type sqlConfig struct {
	ClientAlias string `json:"clientAlias"` //连接别名
	DriverName  string `json:"driverName"`  //数据库名称
	DataSource  string `json:"dataSource"`  //数据库连接信息
}

// GetSQLClient 获取指定名称的数据源
func GetSQLClient(alias string) *sql.DB {
	client, ok := sqlClients[alias]
	if !ok {
		panic("指定名称的数据源不存在")
	}
	InfoLogger.Printf("获取名称为%s的数据源", alias)
	return client
}
