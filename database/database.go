package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	db := "root:@tcp(127.0.0.1:3306)/go_fiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(db), &gorm.Config{})

	if err != nil {
		panic("Can't connect to database")
	}

	fmt.Println("Database Connected")

	DB = database
}
