package entry

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-redis/redis"
	//只引用包中的init函数
	_ "github.com/go-sql-driver/mysql"
)

// MysqlDB 初始化一个mysql数据库连接池
var MysqlDB *sql.DB

// RedisClient 初始化一个redis客户端
var RedisClient redis.UniversalClient

func instanceMysqlDB(userName, password, host, dataBase string, args ...string) *sql.DB {
	var param = "?"
	for index, arg := range args {
		if index == 0 {
			param += arg
		}
		param = param + "&" + arg
	}
	db, err := sql.Open("mysql", userName+":"+password+"@tcp("+host+")/"+dataBase+param)
	if err != nil {
		log.Printf("初始化数据库连接出错！错误信息：%s", err.Error())
		os.Exit(2)
	}
	return db
}
func instanceRediClient(host []string) redis.UniversalClient {
	return redis.NewClusterClient(&redis.ClusterOptions{Addrs: host})
}
func init() {
	MysqlDB = instanceMysqlDB("root", "Scfy774250.", "127.0.0.1:3306", "scfy_im", "charset=UTF8")
	RedisClient = instanceRediClient([]string{"", "", "", ""})
}
