package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/Scfy-Code/IM/pkg/client"
	"github.com/Scfy-Code/IM/pkg/view"
	"github.com/go-redis/redis"
)

// APP 应用的配置信息
var APP *application = new(application)

type application struct {
	TemplateDir    string                  `json:"templateDir"`    //视图目录
	StaticDir      string                  `json:"staticDir"`      //静态文件目录
	UploadDir      string                  `json:"uploadDir"`      //上传文件目录
	LoggerDir      string                  `json:"loggerDir"`      //日志目录
	DriverName     string                  `json:"driverName"`     //数据库
	DataSourceName string                  `json:"dataSourceName"` //数据源信息
	RedisOptions   *redis.UniversalOptions `json:"redisOptions"`   //缓存配置
}

func init() {
	data, err := ioutil.ReadFile("../web/application.json")
	if err != nil {
		log.Printf("读取配置文件出错！错误信息：%s", err.Error())
		os.Exit(2)
	}
	err0 := json.Unmarshal(data, APP)
	if err0 != nil {
		log.Printf("解析配置文件出错！错误信息：%s", err0.Error())
		os.Exit(2)
	}
	view.RegistTemplateDir(APP.TemplateDir)
	client.RegistClient(APP.DriverName, APP.DataSourceName)
}
