package mapper

import (
	"github.com/Scfy-Code/IM/app/livechat/entity"
	"github.com/Scfy-Code/IM/sys"
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
}

func newTeamMapperImpl() TeamMapper {
	return &teamMapperImpl{}
}

func (tmi teamMapperImpl) CreateTeam(teamID string) bool {
	return false
}
func (tmi teamMapperImpl) DeleteTeam(bindID string) bool {
	var sql = "delete from team_user where id = ?"
	stmt, err0 := sys.ReturnSQLClient("US").Prepare(sql)
	if err0 != nil {
		return false
	}
	_, err1 := stmt.Exec(bindID)
	if err1 != nil {
		return false
	}
	return true
}
func (tmi teamMapperImpl) UpdateTeam(teamInfo map[string]interface{}) bool {
	return false
}
func (tmi teamMapperImpl) SelectTeam(talkerID string) map[string]interface{} {
	return nil
}
func (tmi teamMapperImpl) SelectTeams(selfID string) []map[string]interface{} {
	var (
		sql = `SELECT 
					tu.id AS bindID, t.id, t.name, t.notice, t.avatar
				FROM
					team_user tu
						LEFT JOIN
					team t ON tu.teamID = t.id
				WHERE
				tu.memberID = ?`
		result []map[string]interface{}
	)
	rows, err0 := sys.ReturnSQLClient("US").Query(sql, selfID)
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
func NewTeamMapper(mapperName string) TeamMapper {
	switch mapperName {
	case "teamMapper":
		return newTeamMapperImpl()
	default:
		return newTeamMapperImpl()
	}
}
