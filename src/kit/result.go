package kit

import (
	"database/sql"
)

// RowsToMap 将查询的结果集转换成字典切片
func RowsToMap(rows *sql.Rows) []map[string]interface{} {
	columnsType, err1 := rows.ColumnTypes()
	if err1 != nil {
		return nil
	}
	result := new([]map[string]interface{})
	for rows.Next() {
		entry := make(map[string]interface{}, len(columnsType))
		scanner := make([]interface{}, len(columnsType))
		for i := 0; i < len(columnsType); i++ {
			var data interface{}
			switch columnsType[i].DatabaseTypeName() {
			case "VARCHAR":
				data = new(string)
			case "TEXT":
				data = new(int)
			case "NVARCHAR":
				data = new(string)
			case "DECIMAL":
				data = new(float64)
			case "BOOL":
				data = new(bool)
			case "INT":
				data = new(int32)
			case "BIGINT":
				data = new(int64)
			}
			entry[columnsType[i].Name()] = data
			scanner[i] = data
		}
		rows.Scan(scanner...)
		*result = append(*result, entry)
	}
	return *result
}
