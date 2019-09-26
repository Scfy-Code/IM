package util

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
)

// CheckEmail 验证邮箱
func CheckEmail(email string) bool {
	result, err := regexp.MatchString("", email)
	if err != nil {
		return false
	}
	return result
}

// CheckPassword 验证密码
func CheckPassword(password string) bool {
	result, err := regexp.MatchString("", password)
	if err != nil {
		return false
	}
	return result
}
func columnToMap(columnTypes []*sql.ColumnType) (columnSclise []interface{}, columnMap map[string]interface{}) {
	columnSclise = make([]interface{}, len(columnTypes))
	columnMap = make(map[string]interface{}, len(columnTypes))
	for index, columnType := range columnTypes {
		var column interface{}
		column = reflect.New(columnType.ScanType()).Interface()
		columnMap[columnType.Name()] = column
		columnSclise[index] = column
	}
	return
}

// RowsToMap 将值输入map中
func RowsToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()
	columnTypes, err1 := rows.ColumnTypes()
	if err1 != nil {
		return nil, fmt.Errorf("解析结果集参数类型列表！%w", err1)
	}
	var sm []map[string]interface{}
	var err error
	for rows.Next() {
		row, rs := columnToMap(columnTypes)
		err2 := rows.Scan(row...)
		if err2 != nil {
			err = fmt.Errorf("解析结果集出现错误！%w", err2)
			continue
		}
		sm = append(sm, rs)
	}
	return sm, err
}
