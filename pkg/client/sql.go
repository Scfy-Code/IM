package client

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // 只使用初始化方法
)

// SQLClient sql客户端
var SQLClient *sql.DB

// RegistClient 注册一个信息
func RegistClient(driverName, dataSourceName string) {
	db, err0 := sql.Open(driverName, dataSourceName)
	if err0 != nil {
		os.Exit(1)
	}
	err1 := db.Ping()
	if err1 != nil {
		os.Exit(1)
	}
	SQLClient = db
}
