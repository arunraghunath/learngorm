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

func TestSaveAndFirst(t *testing.T) {
	db := getDB()
	newUser := &User{
		Name: "Arun",
	}
	orm := db.Save(newUser)

	queryUser := &User{}
	orm.First(queryUser)
	fmt.Println("Printing user details as -->", queryUser.Name)
}
