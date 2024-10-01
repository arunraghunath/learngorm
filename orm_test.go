package lorm

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
}

func getDB() DB {
	db, _ := Open("postgres", "user=arun.raghunath dbname=arun.raghunath sslmode=disable")
	return db
}

func TestSave(t *testing.T) {
	db := getDB()
	user := &User{
		Name: "Arun",
	}
	orm := db.Save(user)
	fmt.Println(orm.Sql)
	fmt.Println(orm.SqlVars...)

}
