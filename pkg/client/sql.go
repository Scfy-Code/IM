package client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// SQLClient sql客户端
var SQLClient *sql.DB

// RegistClient 注册一个信息
func RegistClient(driverName, dataSourceName string) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		os.Exit(1)
	}
	SQLClient = db
}
