package controller

import (
	"example/firstApp/services"
	"example/firstApp/validator"

	"net/http"

	"github.com/gin-gonic/gin"
)

// get all users
func GetUsers(c *gin.Context) {

	users, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// create new user
func CreateUser(c *gin.Context) {
	// validation
	var userInput validator.UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store user
	user, err := services.CreateUser(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// get user by id
func GetUser(c *gin.Context) {
	// calidate id - check if integer
	id, err := services.CheckID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter valide number id"})
		return
	}

	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// update user
func UpdateUser(c *gin.Context) {
	// calidate id - check if integer
	id, err := services.CheckID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter valide number id"})
		return
	}

	// validation
	var userInput validator.UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update user
	user, err := services.UpdateUser(id, userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "User Updated Successfully", "data": user})
}

// delete user
func DeleteUser(c *gin.Context) {
	// calidate id - check if integer
	id, err := services.CheckID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter valide number id"})
		return
	}

	Err := services.DeleteUser(id)
	if Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully"})
}
