package app

import "database/sql"

var (
	// MysqlClient mysql数据库客户端
	MysqlClient *sql.DB
)

//mysql配置
type mysql struct {
	DriverName string `json:"driverName"` //关系型数据库名称
	UserName   string `json:"userName"`   //用户名
	Password   string `json:"password"`   //密码
	Host       string `json:"host"`       //主机
	DataBase   string `json:"dataBase"`   //数据库
	Charset    string `json:"charset"`    //编码
}
