package entity

// Talker 好友结构体
type talker struct {
	bindID         int64
	talkerID       int64
	talkerNickName string
	talkerAvatar   string
	talkerSign     string
	status         bool
}

func (t *talker) GetFields() []interface{} {
	return []interface{}{
		&t.bindID,
		&t.talkerID,
		&t.talkerNickName,
		&t.talkerAvatar,
		&t.talkerSign,
		&t.status,
	}
}
func (t *talker) EntityToMap() map[string]interface{} {
	return map[string]interface{}{
		"bindID":         t.bindID,
		"talkerID":       t.talkerID,
		"talkerNickName": t.talkerNickName,
		"talkerAvatar":   t.talkerAvatar,
		"talkerSign":     t.talkerSign,
		"status":         t.status,
	}
}

// NewTalker 创建一个好友实体对象
func NewTalker() Entity {
	return &talker{}
}
