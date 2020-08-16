package goft

import (
	"database/sql"
	"fmt"
	"github.com/shenyisyn/goft-ioc"
)

type GPAUtil struct {
	GDB GPA `inject:"-"`
}

func NewGPAUtil() *GPAUtil {
	return &GPAUtil{}
}

type GPA interface {
	DB() *sql.DB
}

func DBMap(columns []string, rows *sql.Rows) ([]interface{}, error) {
	allRows := make([]interface{}, 0) //所有行  大切片
	for rows.Next() {
		oneRow := make([]interface{}, len(columns)) //定义一行切片
		scanRow := make([]interface{}, len(columns))
		fieldMap := make(map[string]interface{})
		for i, _ := range oneRow {
			scanRow[i] = &oneRow[i]
		}
		err := rows.Scan(scanRow...)
		if err != nil {
			return nil, err
		}
		for i, val := range oneRow {
			v, ok := val.([]byte) //断言
			if ok {
				fieldMap[columns[i]] = string(v)
			}
		}
		allRows = append(allRows, fieldMap)
	}
	return allRows, nil
}

func queryForMapsByInterface(query Query) (interface{}, error) {
	ret, err := queryForMaps(query.Sql(), query.Mapping(), query.Args()...)
	if err != nil {
		panic(err)
	}
	if query.First() && ret != nil && len(ret) > 0 {
		return ret[0], nil
	}
	return ret, nil
}
func queryForMaps(sql string, mapping map[string]string, args ...interface{}) ([]interface{}, error) {
	gpa_bean := Injector.BeanFactory.Get((*GPAUtil)(nil)).(*GPAUtil)
	if gpa_bean.GDB == nil {
		panic("found no GPA-Object")
	}
	stmt, err := gpa_bean.GDB.DB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		panic(fmt.Errorf("sql-error:%s", err.Error()))
	}
	rows, err := stmt.Query(args...)
	defer rows.Close()
	if err != nil {
		panic(fmt.Errorf("sqlquery-error:%s", err.Error()))
	}
	cols, err := rows.Columns()
	if err != nil {
		panic(fmt.Errorf("sqlcolumn-error:%s", err.Error()))
	}
	if mapping != nil && len(mapping) > 0 {
		newCols := []string{}
		for _, col := range cols {
			if v, ok := mapping[col]; ok {
				newCols = append(newCols, v)
			} else {
				newCols = append(newCols, col)
			}
		}
		return DBMap(newCols, rows)
	} else {
		return DBMap(cols, rows)
	}

}
