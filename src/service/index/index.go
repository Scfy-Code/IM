package index

import (
	"database/sql"

	"github.com/Scfy-Code/scfy-im/util"

	"github.com/Scfy-Code/scfy-im/database"
)

type IndexService struct {
	db *sql.DB
}

func NewIndexService() *IndexService {
	return &IndexService{database.MysqlClient}
}
func (is IndexService) SelectFriends(id string) []map[string]interface{} {
	rows, err := is.db.Query("query", id)
	if err != nil {
		return nil
	}
	return util.RowsToMap(rows)
}
func (is IndexService) SelectGroups(id string) []map[string]interface{} {
	rows, err := is.db.Query("query", id)
	if err != nil {
		return nil
	}
	return util.RowsToMap(rows)
}
