package entity

type account struct {
	email    string
	password string
}

// NewAccount 创建账号对象
func NewAccount() Entity {
	return &account{}
}
func (a *account) GetFields() []interface{} {
	return []interface{}{&a.email, &a.password}
}
func (a *account) EntityToMap() map[string]interface{} {
	return map[string]interface{}{"email": a.email, "password": a.password}
}
