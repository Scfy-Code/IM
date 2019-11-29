package sys

import "github.com/go-redis/redis"

type application struct {
	Port           string                  `json:"port"`           //监听端口
	RuntimeEnv     string                  `json:"runtimeEnv"`     //运行环境
	TemplateDir    string                  `json:"templateDir"`    //视图目录
	StaticDir      string                  `json:"staticDir"`      //静态文件目录
	UploadDir      string                  `json:"uploadDir"`      //上传文件目录
	LoggerDir      string                  `json:"loggerDir"`      //日志目录
	DriverName     string                  `json:"driverName"`     //数据库
	DataSourceName string                  `json:"dataSourceName"` //数据源信息
	RedisOptions   *redis.UniversalOptions `json:"redisOptions"`   //缓存配置
}
