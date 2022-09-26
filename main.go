package main

import (
	"example/firstApp/api/course"
	"example/firstApp/api/material"
	"example/firstApp/api/user"
	"example/firstApp/database"
	"example/firstApp/route"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// connect to data base
var db *gorm.DB = database.ConnectToSQL()

func main() {
	// close connection
	defer database.CloseConnection(db)

	// need to handle
	db.AutoMigrate(&user.User{}, &course.Course{}, &material.Material{})

	// initialize router
	router := gin.New()

	// route
	routes.MainRoute(router.Group(""))

	// connect to and run the server
	router.Run(os.Getenv("LOCAL_HOST") + ":" + os.Getenv("SERVER_PORT"))
}
