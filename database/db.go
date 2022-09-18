package database

import (
	"example/firstApp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	
	dsn := "root:@tcp(127.0.0.1:3306)/go_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Faild to connect database")
	}

	db.AutoMigrate(&userModel.User{})

	DB = db
}
