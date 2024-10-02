package lorm

import (
	"fmt"
	"testing"
)

type User struct {
	Username string
}

func getDB() DB {
	db, _ := Open("postgres", "user=arun.raghunath dbname=arun.raghunath sslmode=disable")
	return db
}

func TestSaveAndFirst(t *testing.T) {
	db := getDB()
	newUser := &User{
		Username: "Arun",
	}
	db.Save(newUser)

	queryUser := &User{}
	db.First(queryUser)
	fmt.Println("Printing user details as -->", queryUser.Username)

	queryUsers := []User{}
	db.First(&queryUsers)
	for _, v := range queryUsers {
		fmt.Println("Printing slice of users ", v.Username)
	}
}

func TestDeleteAndFirst(t *testing.T) {
	db := getDB()
	user := User{}
	db.Where("Username = ?", "Arun").Delete(user)
	fmt.Println("Printing user deleted")
	getUser := &User{}
	orm := db.First(getUser)
	fmt.Println(orm.Error)
}

func TestWhere(t *testing.T) {
	db := getDB()
	user := &User{Username: "Arun"}
	db.Save(user)
	whereUser := &User{}
	db.Where("Name = ?", "Arun").First(whereUser)
	fmt.Println("Queried user detail is --> ", whereUser.Username)
}
