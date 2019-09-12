package index

import (
	"database/sql"

	"github.com/Scfy-Code/scfy-im/kit"

	"github.com/Scfy-Code/scfy-im/app"
)

type IndexService struct {
	db *sql.DB
}

func NewIndexService() *IndexService {
	return &IndexService{app.MysqlClient}
}
func (is IndexService) SelectFriends(id string) []map[string]interface{} {
	rows, err := is.db.Query("select user1 , user2 from nexus where user1 = ?", id)
	if err != nil {
		return nil
	}
	return kit.RowsToMap(rows)
}
func (is IndexService) SelectGroups(id string) []map[string]interface{} {
	rows, err := is.db.Query("query", id)
	if err != nil {
		return nil
	}
	return kit.RowsToMap(rows)
}
