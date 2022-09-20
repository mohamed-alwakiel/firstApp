package database

import (
	"fmt"
	"os"

	"example/firstApp/user"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectToSQL() *gorm.DB {

	// check if env file is exist or not
	err_env := godotenv.Load()
	if err_env != nil {
		panic("Faild to load (env) file !")
	}

	// instance of DbConfig struct with info form env file
	db_config := DbConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_config.User,
		db_config.Password,
		db_config.Host,
		db_config.Port,
		db_config.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Faild to connect to database !")
	}

	db.AutoMigrate(&user.User{})
	DB = db

	return db
}

func CloseConnection(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Faild to get connection !")
	}
	conn.Close()
}
