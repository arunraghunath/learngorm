package lorm

import (
	"fmt"
	"reflect"
	"strings"
)

func (orm *Orm) preparePlan(value interface{}, operation string) {
	orm.setUser()
	switch operation {
	case "save":
		orm.saveSql(value)
	case "query":
		orm.querySql(value)
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

func (orm *Orm) query(value interface{}) {
	rows, err := orm.db.Query(orm.Sql)
	orm.Error = err
	fields := reflect.TypeOf(value).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fmt.Println("Printing Struct field name ", fields.Field(i).Name)
	}
	for rows.Next() {
		dest := reflect.ValueOf(value).Elem()
		fmt.Printf("%+v", dest)
		columns, _ := rows.Columns()
		var values []interface{}
		for _, v := range columns {
			fmt.Println("Column name is ", v)
			values = append(values, dest.FieldByName(v).Addr().Interface())

		}
		fmt.Println(values)
		orm.Error = rows.Scan(values...)

	}
}
