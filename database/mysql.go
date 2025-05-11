package database

import (
	"gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/rest_api_go"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(models.Product{}, &models.User{})
	DB = database
}
