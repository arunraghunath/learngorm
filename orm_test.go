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
	db.Save(newUser)

	queryUser := &User{}
	db.First(queryUser)
	fmt.Println("Printing user details as -->", queryUser.Name)

	queryUsers := []User{}
	db.First(&queryUsers)
	for _, v := range queryUsers {
		fmt.Println("Printing slice of users ", v.Name)
	}
}

func TestDeleteAndFirst(t *testing.T) {
	db := getDB()
	user := User{}
	db.Delete(user)
	fmt.Println("Printing user deleted")
	getUser := User{}
	orm := db.First(getUser)
	fmt.Println(orm.Error)

}
