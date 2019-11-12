package mapper

import (
	"database/sql"

	"github.com/Scfy-Code/IM/app/entity"
	"github.com/Scfy-Code/IM/pkg/client"
	_ "github.com/Scfy-Code/IM/pkg/conf" //只使用初始化方法
)

// IndexMapper 首页数据映射接口
type IndexMapper interface {
	GetTalkerList(selfID string) []map[string]interface{}
	GetTeamList(selfID string) []map[string]interface{}
}
type indexMapperImpl struct {
	sqlClient *sql.DB
}

func (imi indexMapperImpl) GetTalkerList(selfID string) []map[string]interface{} {
	var (
		talkerList []map[string]interface{}
	)
	rows, err0 := imi.sqlClient.Query("SELECT l.friendID AS id, IF(l.nickName1 IS NULL, s.nickName, l.nickName1) AS nickName, s.avatar, s.sign, FALSE AS status FROM talker_talker l LEFT JOIN talker s ON l.friendID = s.id WHERE selfID = ? UNION ALL SELECT l.selfID AS id, IF(l.nickName2 IS NULL, s.nickName, l.nickName2) AS nickName, s.avatar, s.sign, FALSE AS status FROM talker_talker l LEFT JOIN talker s ON l.friendID = s.id WHERE friendID = ? ", selfID, selfID)
	if err0 != nil {

	}
	for rows.Next() {
		var (
			talker = entity.NewTalker()
		)
		err1 := rows.Scan(talker.GetFields()...)
		if err1 != nil {
			continue
		}
		talkerList = append(talkerList, talker.EntityToMap())
	}
	return talkerList
}
func (imi indexMapperImpl) GetTeamList(selfID string) []map[string]interface{} {
	var (
		teamList []map[string]interface{}
	)
	rows, err0 := imi.sqlClient.Query("SELECT t.id, t.name, t.notice,t.avatar FROM team_talker tt LEFT JOIN team t ON tt.teamID = t.id WHERE tt.talkerID =?", selfID)
	if err0 != nil {

	}
	for rows.Next() {
		team := entity.NewTeam()
		err1 := rows.Scan(team.GetFields()...)
		if err1 != nil {
			continue
		}
		teamList = append(teamList, team.EntityToMap())
	}
	return teamList
}

// NewIndexMapper 创建首页数据映射实例
func NewIndexMapper() IndexMapper {
	return indexMapperImpl{
		client.SQLClient,
	}
}
