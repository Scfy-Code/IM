package mapper

// AccountMapper 数据库交互接口
type AccountMapper interface {
	SelectAccount(email, password string) map[string]interface{}
}
