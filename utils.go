package lorm

import (
	"fmt"
	"reflect"
	"strings"
)

func modelValues(m interface{}) (columns []string, values []interface{}) {
	typ := reflect.TypeOf(m).Elem()
	fmt.Println(typ)
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			columns = append(columns, strings.ToLower(p.Name))
			value := reflect.ValueOf(m).Elem().FieldByName(p.Name)
			values = append(values, value.Interface())
		}
	}
	return
}

func quoteMap(columns []string) (results []string) {
	for _, v := range columns {
		results = append(results, "\""+v+"\"")
	}
	return
}

func valuesToBindVar(values []interface{}) string {
	var sqls []string
	for index, _ := range values {
		sqls = append(sqls, fmt.Sprintf("$%d", index+1))
	}
	return strings.Join(sqls, ",")
}
