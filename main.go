package main

import (
	// "net/http"
	"example/firstApp/controllers"
	"example/firstApp/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize router
	r := gin.Default()

	// connect to data base
	database.ConnectDatabase()

	// routes
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
