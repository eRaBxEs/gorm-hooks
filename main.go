package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {

	var err error

	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Automigrate the models to create the users table
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to auto-migrate ")
	}

	// create a new user
	user := User{Name: "Johnny", Email: "johnny@gmail"}
	result := db.Create(&user)
	if result.Error != nil {
		panic("failed to create user")
	}

}
