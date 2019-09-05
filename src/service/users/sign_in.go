package users

import (
	"database/sql"
	"log"

	"github.com/Scfy-Code/scfy-im/entry"
	_ "github.com/go-sql-driver/mysql"
)

var dbs *sql.DB

// UserService 数据交互结构体
type UserService struct {
	Conn *sql.DB
}

// Login 登录方法
func (us UserService) Login(email, password string) *entry.UserEntry {
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

// NewUserService 新建一个数据交互结构体
func NewUserService() *UserService {
	if dbs == nil {
		return nil
	} else {
		return &UserService{dbs}
	}
}
func init() {
	db, err := sql.Open("mysql", "root:Scfy774250.@tcp(127.0.0.1:3306)/scfy_im?charset=UTF8")
	if err == nil {
		dbs = db
	} else {
		dbs = nil
	}
}
