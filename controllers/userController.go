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

	// return response
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

	// return response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// get user by id
func GetUser(c *gin.Context) {
	var user userModel.User

	// check if exist
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found !"})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// update user
func UpdateUser(c *gin.Context) {
	var user userModel.User

	// check if exists
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found !"})
		return
	}

	// validation
	var userInput userModel.UpdateUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update user
	database.DB.Model(&user).Updates(userInput)

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "User Updated Successfully", "data": user})
}

// delete user
func DeleteUser(c *gin.Context) {
	var user userModel.User

	// check if exist
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found !"})
		return
	}

	// delete user
	database.DB.Delete(&user)

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully"})
}
