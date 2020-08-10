package goft

import (
	"database/sql"
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

func queryForMaps(sql string) ([]interface{}, error) {
	gpa_bean := Injector.BeanFactory.Get((*GPAUtil)(nil)).(*GPAUtil)
	if gpa_bean.GDB == nil {
		panic("found no GPA-Object")
	}

	rows, err := gpa_bean.GDB.DB().Query(sql)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	cols, _ := rows.Columns()
	return DBMap(cols, rows)
}
