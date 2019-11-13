package mapper

import (
	"database/sql"

	"github.com/Scfy-Code/IM/app/entity"
	"github.com/Scfy-Code/IM/pkg/client"
	_ "github.com/Scfy-Code/IM/pkg/conf" //只使用初始化方法
)

// TeamMapper 群组数据层映层射接口
type TeamMapper interface {
	CreateTeam(teamID string) bool
	DeleteTeam(teamID string) bool
	UpdateTeam(teamInfo map[string]interface{}) bool
	SelectTeam(talkerID string) map[string]interface{}
	SelectTeams(selfID string) []map[string]interface{}
}
type teamMapperImpl struct {
	sqlClient *sql.DB
}

func (tmi teamMapperImpl) CreateTeam(teamID string) bool {
	return false
}
func (tmi teamMapperImpl) DeleteTeam(teamID string) bool {
	return false
}
func (tmi teamMapperImpl) UpdateTeam(teamInfo map[string]interface{}) bool {
	return false
}
func (tmi teamMapperImpl) SelectTeam(talkerID string) map[string]interface{} {
	return nil
}
func (tmi teamMapperImpl) SelectTeams(selfID string) []map[string]interface{} {
	var (
		sql    = "SELECT tt.id as bindID, t.id, t.name, t.notice,t.avatar FROM team_talker tt LEFT JOIN team t ON tt.teamID = t.id WHERE tt.talkerID =?"
		result []map[string]interface{}
	)
	rows, err0 := tmi.sqlClient.Query(sql, selfID)
	if err0 != nil {

	}
	for rows.Next() {
		team := entity.NewTeam()
		err1 := rows.Scan(team.GetFields()...)
		if err1 != nil {
			continue
		}
		result = append(result, team.EntityToMap())
	}
	return result
}

// NewTeamMapper 新建群组数据映射层对象
func NewTeamMapper() TeamMapper {
	return teamMapperImpl{
		client.SQLClient,
	}
}
