package users

import (
	"database/sql"

	"github.com/Scfy-Code/scfy-im/database"
	"github.com/Scfy-Code/scfy-im/log"
	"github.com/Scfy-Code/scfy-im/util"
)

// UserService 对外开放一个对象便于操作数据
var UserService = &userService{database.MysqlDB}

// userService 数据交互结构体
type userService struct {
	Conn *sql.DB
}

// Login 登录方法
func (us userService) Login(email, password string) map[string]interface{} {
	row, err := us.Conn.Query("select id,remarkname,avatar,email,ignature from user where email=? and password=?", email, password)
	if err != nil {
		log.WarnLog(err.Error())
		return nil
	}
	if row.Next() {
		user := util.RowsToMap(row)[0]
		return user
	}
	return nil
}
