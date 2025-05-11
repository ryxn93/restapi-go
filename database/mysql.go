package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	database, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306/rest_api_go)"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate()

	DB = database
}
