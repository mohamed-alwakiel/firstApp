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
	var userInput userModel.CreateUserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store user
	user := userModel.User{
		Name:  userInput.Name,
		Email: userInput.Email,
		Age:   20,
	}

	database.DB.Create(&user)

	// response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// get user by id
func GetUser(c *gin.Context) {
	var user userModel.User

	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found !"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
