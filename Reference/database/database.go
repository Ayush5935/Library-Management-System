package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	host := "localhost"
	port := "5432"
	database := "TaskDb"
	user_name := "postgres"
	password := ""

	db, err := gorm.Open(postgres.Open("postgres://" + user_name + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"))
	if err != nil {
		fmt.Println(err, " Database Connection Failed")
		log.Fatal("connection error: ", err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
	// db.AutoMigrate(&User{}, &Library{}, &BookInventory{}, &RequestEvents{}, &IssueRegistry{})
	DB = db
}
