package users

import (
	"database/sql"
	"log"

	"github.com/Scfy-Code/scfy-im/entry"
)

// UserService 对外开放一个对象便于操作数据
var UserService = &userService{entry.MysqlDB}

// userService 数据交互结构体
type userService struct {
	Conn *sql.DB
}

// Login 登录方法
func (us userService) Login(email, password string) *entry.UserEntry {
	row, err := us.Conn.Query("select id,account,remarkname,avatar,email from user where email=? and password=?", email, password)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	var user = &entry.UserEntry{}
	if row.Next() {
		row.Scan(user.ID, user.Account, user.RemarkName, user.Avatar, user.Email)
	}
	return user
}
