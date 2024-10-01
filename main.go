package lorm

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	Db *sql.DB
}

func Open(driver, source string) (gdb DB, err error) {
	gdb.Db, err = sql.Open(driver, source)
	return
}

func (gdb *DB) initializeOrm() *Orm {
	orm := &Orm{
		db: gdb.Db,
	}
	return orm
}
func (gdb *DB) Save(value interface{}) *Orm {
	orm := gdb.initializeOrm()
	orm = orm.Save(value)
	return orm
}
