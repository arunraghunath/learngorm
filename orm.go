package lorm

import (
	"database/sql"
	"fmt"
)

type Orm struct {
	TableName   string
	PrimaryKey  string
	SqlResult   sql.Result
	Error       error
	Sql         string
	SqlVars     []interface{}
	db          *sql.DB
	whereClause []interface{}
	selectStr   string
	orderStr    string
	operation   string
}

func (orm *Orm) setUser() {
	orm.TableName = "users"
	orm.PrimaryKey = "id"
}

func (orm *Orm) Save(value interface{}) *Orm {
	orm.preparePlan(value, "save")
	orm.Execute()
	return orm
}

func (orm *Orm) First(out interface{}) *Orm {
	orm.preparePlan(out, "query")
	orm.query(out)
	fmt.Printf("Just finished First %+v\n", out)
	return orm
}

func (orm *Orm) Execute() *Orm {
	orm.SqlResult, orm.Error = orm.db.Exec(orm.Sql, orm.SqlVars...)
	return orm
}

func (orm *Orm) Delete(value interface{}) *Orm {
	orm.preparePlan(value, "delete")
	orm.Execute()
	return orm
}
