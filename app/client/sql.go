package client

import (
	"database/sql"

	"github.com/Scfy-Code/IM/app"
	"github.com/Scfy-Code/IM/app/mylog"
	_ "github.com/go-sql-driver/mysql" // 只使用初始化方法
)

var (
	// SQLClient sql客户端
	SQLClient *sql.DB
	logger    = mylog.NewWarnLogger()
)

func init() {
	db, err0 := sql.Open(pkg.APP.DriverName, pkg.APP.DataSourceName)
	if err0 != nil {
		logger.Fatalf("创建SQL客户端出错！错误信息：%s", err0.Error())
	}
	err1 := db.Ping()
	if err1 != nil {
		logger.Fatalf("连接SQL客户端出错！错误信息：%s", err1.Error())
	}
	SQLClient = db
}
