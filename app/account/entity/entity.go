package entity

// Entity 实体类接口
type Entity interface {
	GetFields() []interface{}
	EntityToMap() map[string]interface{}
}
