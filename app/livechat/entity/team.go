package entity

type team struct {
	bindID int64
	id     int64
	Name   string
	Notice string
	Avatar string
}

func (t *team) GetFields() []interface{} {
	return []interface{}{
		&t.bindID,
		&t.id,
		&t.Name,
		&t.Notice,
		&t.Avatar,
	}
}
func (t *team) EntityToMap() map[string]interface{} {
	return map[string]interface{}{
		"bindID":     t.bindID,
		"teamID":     t.id,
		"teamName":   t.Name,
		"teamNotice": t.Notice,
		"teamAvatar": t.Avatar,
	}
}

// NewTeam 创建群组实体对象
func NewTeam() Entity {
	return &team{}
}
