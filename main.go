package main

import (
	"fmt"
	// "net/http"

	"github.com/mohamed-alwakiel/firstApp/models"

)

// type User struct {
// 	ID    uint    `json:"id" gorm:"primary_key"`
// 	Name  string  `json:"user_name"`
// 	Email *string `json:"email"`
// 	Age   uint    `json:"user_age"`
// }

func main() {
	fmt.Println("Hello Wiko")
	models.HelloWiko()
	// initialize router
	// r := gin.Default()

	// connect to data base
	// models.ConnectDatabase()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	// })

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.Run()
}
