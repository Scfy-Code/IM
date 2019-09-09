package database

import (
	"database/sql"
	"os"

	"github.com/Scfy-Code/scfy-im/config"
	"github.com/Scfy-Code/scfy-im/log"
	"github.com/go-redis/redis"

	//只引用包中的init函数
	_ "github.com/go-sql-driver/mysql"
)

var (
	// MysqlClient mysql客户端
	MysqlClient = newMysqlClient()
	// redisClient redis客户端
	redisClient = newRedisClient()
)

// NewMysqlClient 创建一个mysql客户端
func newMysqlClient() *sql.DB {
	db, err := sql.Open(config.APPCFG.DataBaseCfg.DriverName, config.APPCFG.DataBaseCfg.DataSourceName)
	if err != nil {
		log.WarnLog(err.Error())
		os.Exit(2)
	}
	return db
}

// NewRedisClient 创建一个redis客户端
func newRedisClient() redis.UniversalClient {
	return redis.NewClusterClient(&redis.ClusterOptions{Addrs: config.APPCFG.RedisCfg.Addrs})
}
