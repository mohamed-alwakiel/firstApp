package main

import (
	"example/firstApp/database"
	"example/firstApp/route"
	"example/firstApp/user"
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
	db.AutoMigrate(&user.User{})

	// initialize router
	router := gin.New()

	// route
	routes.MainRoute(router.Group(""))

	// connect to and run the server
	router.Run(os.Getenv("LOCAL_HOST") + ":" + os.Getenv("SERVER_PORT"))
}
