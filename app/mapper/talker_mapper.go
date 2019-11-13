package mapper

import (
	"database/sql"

	"github.com/Scfy-Code/IM/app/entity"
	"github.com/Scfy-Code/IM/pkg/client"
	_ "github.com/Scfy-Code/IM/pkg/conf" //只使用初始化方法
)

// TalkerMapper 好友数据库映射对象
type TalkerMapper interface {
	CreateTalker(talkerID string) bool
	DeleteTalker(bindID string) bool
	UpdateTalker(talkerInfo map[string]interface{}) bool
	SelectTalker(talkerID string) map[string]interface{}
	SelectTalkers(selfID string) []map[string]interface{}
}
type talkerMapperImpl struct {
	sqlClient *sql.DB
}

func (tmi talkerMapperImpl) CreateTalker(talkerID string) bool {
	return false
}
func (tmi talkerMapperImpl) DeleteTalker(bindID string) bool {
	var (
		sql = "DELETE FROM talker_talker WHERE id = ?"
	)
	stmt, err0 := tmi.sqlClient.Prepare(sql)
	if err0 != nil {
		return false
	}
	_, err1 := stmt.Exec(bindID)
	if err1 != nil {
		return false
	}
	return true
}
func (tmi talkerMapperImpl) UpdateTalker(talkerInfo map[string]interface{}) bool {
	return false
}
func (tmi talkerMapperImpl) SelectTalker(TalkerID string) map[string]interface{} {
	var (
		sql    = ""
		result map[string]interface{}
	)
	rows, err0 := tmi.sqlClient.Query(sql, TalkerID, TalkerID)
	if err0 != nil {

	}
	if rows.Next() {
		var talker = entity.NewTalker()
		err1 := rows.Scan(talker.GetFields()...)
		if err1 != nil {
			return result
		}
		result = talker.EntityToMap()
	}
	return result
}

func (tmi talkerMapperImpl) SelectTalkers(selfID string) []map[string]interface{} {
	var (
		sql    = "SELECT l.id AS bindID, l.friendID AS id, IF(l.nickName1 IS NULL, s.nickName, l.nickName1) AS nickName, s.avatar, s.sign, FALSE AS status FROM talker_talker l LEFT JOIN talker s ON l.friendID = s.id WHERE selfID = ? UNION ALL SELECT l.id AS bindID, l.selfID AS id, IF(l.nickName2 IS NULL, s.nickName, l.nickName2) AS nickName, s.avatar, s.sign, FALSE AS status FROM talker_talker l LEFT JOIN talker s ON l.friendID = s.id WHERE friendID = ?"
		result []map[string]interface{}
	)
	rows, err0 := tmi.sqlClient.Query(sql, selfID, selfID)
	if err0 != nil {

	}
	for rows.Next() {
		var talker = entity.NewTalker()
		err1 := rows.Scan(talker.GetFields()...)
		if err1 != nil {
			continue
		}
		result = append(result, talker.EntityToMap())
	}
	return result
}

// NewTalkerMapper 新建好友数据映射对象
func NewTalkerMapper() TalkerMapper {
	return talkerMapperImpl{
		client.SQLClient,
	}
}
