package goft

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-ioc"
	"reflect"
	"strconv"
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

func asString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv := reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
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
			fieldMap[columns[i]] = asString(val)
		}
		allRows = append(allRows, fieldMap)
	}
	return allRows, nil
}

func wrapResult(query Query, ret interface{}) interface{} {
	if query.Key() != "" {
		return gin.H{query.Key(): ret}
	}
	return ret
}
func queryForMapsByInterface(query Query) (interface{}, error) {
	ret, err := queryForMaps(query.Sql(), query.Mapping(), query.Args()...)
	if err != nil {
		panic(err)
	}
	if query.First() && ret != nil && len(ret) > 0 {
		return wrapResult(query, ret[0]), nil
	}
	return wrapResult(query, ret), nil
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
