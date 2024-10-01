package lorm

import (
	"fmt"
	"strings"
)

func (orm *Orm) preparePlan(value interface{}, operation string) {
	orm.setUser()
	switch operation {
	case "save":
		orm.saveSql(value)
	}
}

func (orm *Orm) saveSql(value interface{}) {
	columns, values := modelValues(value)
	orm.Sql = fmt.Sprintf("INSERT INTO \"%v\" (%v) VALUES (%v)", orm.TableName,
		strings.Join(quoteMap(columns), ","), valuesToBindVar(values))
	orm.SqlVars = values
}
