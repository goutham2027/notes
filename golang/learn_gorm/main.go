package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=gorm password='' dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// db.SingularTable(true)
	db.DropTable(&User{})
	db.CreateTable(&User{})

	// user := User{
	// 	Username:  "adent",
	// 	FirstName: "Arthur",
	// 	LastName:  "Dent",
	// }

	// db.Create(&user)
	// fmt.Println(user)

	for _, user := range users {
		db.Create(&user)
	}

	// u := User{}
	// db.Last(&u)
	// fmt.Println(u)

	// u := User{Username: "tmacmillan"}
	// db.Where(&u).First(&u)
	// fmt.Println(u)
	// u.LastName = "Beeblebrox"
	// db.Save(&u)

	// fmt.Println(u)

	db.Where(&User{Username: "adent"}).Delete(&User{})

	println("Done")

	// dbase := db.DB()
	// defer dbase.Close()

	// err = dbase.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// println("Connection to database established")
}

type User struct {
	// ID        uint
	gorm.Model
	Username  string
	FirstName string
	LastName  string
}

var users []User = []User{
	User{Username: "adent", FirstName: "Arthur", LastName: "Dent"},
	User{Username: "fprefect", FirstName: "Ford", LastName: "Prefect"},
	User{Username: "tmacmillan", FirstName: "Tricia", LastName: "MacMillan"},
	User{Username: "mrobot", FirstName: "Marvin", LastName: "Robot"},
}
