package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:12345678@tcp(localhost:3306)/go_basics?parseTime=true"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	user := User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
	}

	fmt.Println(user)
}

type User struct {
	gorm.Model
	FirsName string `gorm:"type:VARCHAR(30)"`
	LastName string `gorm:"size:100"`
	Email    string `gorm:"unique"`
}