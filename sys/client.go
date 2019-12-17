package sys

import (
	"database/sql"
	"fmt"
	"reflect"
)

// ReturnToMap 输出map类型
func ReturnToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()
	var (
		resultSclise []map[string]interface{}
	)
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	for i := 0; rows.Next(); i++ {
		var (
			columns []interface{}          = make([]interface{}, len(columnTypes))
			result  map[string]interface{} = make(map[string]interface{}, len(columnTypes))
		)
		for index, columnType := range columnTypes {
			columns[index] = reflect.New(columnType.ScanType()).Interface()
			result[columnType.Name()] = columns[index]
		}
		err = rows.Scan(columns)
		if err != nil {
			err = fmt.Errorf("第%d条数据赋值出错！错误信息%w", i, err)
			continue
		}
		resultSclise = append(resultSclise, result)
	}
	return resultSclise, err
}

// Select 从指定数据源中查询符合条件的结果集
func Select(dataSourceName string, sqlStr string, params ...interface{}) ([]map[string]interface{}, error) {
	rows, err := GetSQLClient(dataSourceName).Query(sqlStr, params...)
	if err != nil {
		return nil, err
	}
	return ReturnToMap(rows)
}

// Update 向指定数据源中新增一条数据
func Update(sql string, params ...interface{}) (int64, error) {
	tx, err := GetSQLClient("dataSourceName").Begin()
	if err != nil {
		return 0, err
	}
	stmt, err := tx.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return 0, nil
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return num, err
}
