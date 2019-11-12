package entity

type team struct {
	id     int64
	Name   string
	Notice string
	Avatar string
}

func (g *team) GetFields() []interface{} {
	return []interface{}{
		&g.id,
		&g.Name,
		&g.Notice,
		&g.Avatar,
	}
}
func (g *team) EntityToMap() map[string]interface{} {
	return map[string]interface{}{
		"teamID":     g.id,
		"teamName":   g.Name,
		"teamNotice": g.Notice,
		"teamAvatar": g.Avatar,
	}
}

// NewTeam 创建群组实体对象
func NewTeam() Entity {
	return &team{}
}
