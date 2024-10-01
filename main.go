package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	Db *sql.DB
}

type User struct {
	Name string
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

func main() {
	db, _ := Open("postgres", "user=arun.raghunath dbname=arun.raghunath sslmode=disable")
	orm := db.Save(&User{Name: "Arun"})
	fmt.Println(orm)
}
