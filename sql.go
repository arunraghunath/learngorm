package lorm

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (orm *Orm) preparePlan(value interface{}, operation string) {
	orm.setUser()
	switch operation {
	case "save":
		orm.saveSql(value)
	case "query":
		orm.querySql(value)
	case "delete":
		orm.deleteSql(value)
	}
}

func (orm *Orm) saveSql(value interface{}) {
	columns, values := modelValues(value)
	orm.Sql = fmt.Sprintf("INSERT INTO \"%v\" (%v) VALUES (%v)", orm.TableName,
		strings.Join(quoteMap(columns), ","), valuesToBindVar(values))
	orm.SqlVars = values
}

func (orm *Orm) querySql(_ interface{}) *Orm {
	orm.Sql = "select * from users limit 1"
	return orm
}

func (orm *Orm) query(out interface{}) {
	var (
		is_slice  bool
		dest_type reflect.Type
	)
	dest_out := reflect.Indirect(reflect.ValueOf(out))
	if kind := dest_out.Kind(); kind == reflect.Slice {
		is_slice = true
		dest_type = dest_out.Type().Elem()
	}

	rows, err := orm.db.Query(orm.Sql)
	orm.Error = err
	for rows.Next() {
		var dest reflect.Value
		if is_slice {
			dest = reflect.New(dest_type).Elem()
		} else {
			dest = reflect.ValueOf(out).Elem()
		}

		fmt.Printf("Printing structure here %+v\n", dest)
		columns, _ := rows.Columns()
		var values []interface{}
		caser := cases.Title(language.English)
		for i, v := range columns {
			fmt.Println("Printing row ", i, v, caser.String(v))
			values = append(values, dest.FieldByName(caser.String(v)).Addr().Interface())
		}
		orm.Error = rows.Scan(values...)
	}
}

func (orm *Orm) deleteSql(value interface{}) {
	orm.Sql = fmt.Sprintf("Delete from %v where %v", orm.TableName, orm.whereSql(value))
	fmt.Println("Delete sql is --> ", orm.Sql)
}

func (orm *Orm) whereSql(value interface{}) (sql string) {
	if len(orm.whereClause) == 0 {
		return
	} else {
		for _, v := range orm.whereClause {
			sql += sql + v["query"].(string)
			args := v["args"].([]interface{})
			for _, arg := range args {
				orm.SqlVars = append(orm.SqlVars, arg.([]interface{})...)
				sql = strings.Replace(sql, "?", "$"+strconv.Itoa(len(orm.SqlVars)), 1)
			}
		}
		return
	}
}
