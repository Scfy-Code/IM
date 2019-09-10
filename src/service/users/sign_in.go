package users

import (
	"database/sql"

	"github.com/Scfy-Code/scfy-im/logger"

	"github.com/Scfy-Code/scfy-im/database"
	"github.com/Scfy-Code/scfy-im/util"
)

// UserService 对外开放一个对象便于操作数据
var UserService = &userService{database.MysqlClient}

// userService 数据交互结构体
type userService struct {
	Conn *sql.DB
}

// Login 登录方法
func (us userService) Login(email, password string) map[string]interface{} {
	row, err := us.Conn.Query("select id,remarkname,avatar,email,signature from user where email=? and password=?", email, password)
	if err != nil {
		logger.WarnPrintf("用户登录出错！错误信息：%s", err.Error())
		return nil
	}
	return util.RowsToMap(row)[0]
}
