package sys

import (
	"database/sql"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" // 只使用初始化方法
)

func init() {
	db, err0 := sql.Open(APP.DriverName, APP.DataSourceName)
	if err0 != nil {
		WarnLogger.Fatalf("创建SQL客户端出错！错误信息：%s", err0.Error())
	}
	err1 := db.Ping()
	if err1 != nil {
		WarnLogger.Fatalf("连接SQL客户端出错！错误信息：%s", err1.Error())
	}
	SQLClient = db
	RedisClient = redis.NewUniversalClient(APP.RedisOptions)
}
