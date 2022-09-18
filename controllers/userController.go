package controllers

import (
	"example/firstApp/database"
	"example/firstApp/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

// get all users
func GetUsers(c *gin.Context) {
	var users []userModel.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// create new user
func CreateUser(c *gin.Context) {
	// validation
	var input userModel.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store user
	user := userModel.User{
		Name:  input.Name,
		Email: input.Email,
		Age:   20,
	}

	database.DB.Create(&user)

	// response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

