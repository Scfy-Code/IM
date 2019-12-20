package mapper

import (
	"github.com/Scfy-Code/IM/app/livechat/entity"
	"github.com/Scfy-Code/IM/sys"
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
}

func newTalkerMapperImpl() TalkerMapper {
	return &talkerMapperImpl{}
}

func (tmi talkerMapperImpl) CreateTalker(talkerID string) bool {
	return false
}
func (tmi talkerMapperImpl) DeleteTalker(bindID string) bool {
	var (
		sql = "DELETE FROM user_user WHERE id = ?"
	)
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
func (tmi talkerMapperImpl) UpdateTalker(talkerInfo map[string]interface{}) bool {
	return false
}
func (tmi talkerMapperImpl) SelectTalker(TalkerID string) map[string]interface{} {
	var (
		sql    = ""
		result map[string]interface{}
	)
	rows, err0 := sys.ReturnSQLClient("US").Query(sql, TalkerID, TalkerID)
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
		sql = `SELECT 
					uu.id AS bindID,
					uu.sponsorID AS id,
					IF(uu.sponsornote IS NULL,
						u.nickName,
						uu.sponsorNote) AS nickName,
					u.avatar,
					u.sign,
					FALSE AS status
				FROM
					user_user uu
						LEFT JOIN
					user u ON uu.sponsorID = u.id
				WHERE
					uu.sponsorID=?
				UNION ALL SELECT 
					uu.id AS bindID,
					uu.sponsorID AS id,
					IF(uu.sponsornote IS NULL,
						u.nickName,
						uu.sponsorNote) AS nickName,
					u.avatar,
					u.sign,
					FALSE AS status
				FROM
					user_user uu
						LEFT JOIN
					user u ON uu.sponsorID = u.id
				WHERE
					uu.receiverID=?`
		result []map[string]interface{}
	)
	rows, err0 := sys.ReturnSQLClient("US").Query(sql, selfID, selfID)
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
func NewTalkerMapper(mapperName string) TalkerMapper {
	switch mapperName {
	case "talkerMapper":
		return newTalkerMapperImpl()
	default:
		return newTalkerMapperImpl()
	}
}
