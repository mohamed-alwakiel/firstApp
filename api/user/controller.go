package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get all users
func GetUsers(c *gin.Context) {
	var filterInput FilterInput
	err := c.ShouldBindJSON(&filterInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get all users
	users, err := GetUsersService(filterInput)
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
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store user
	user, err := CreateUserService(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// get user by id
func GetUser(c *gin.Context) {
	// validate id - check if integer
	id, err := CheckID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter valide number id"})
		return
	}

	user, err := GetUserService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// update user
func UpdateUser(c *gin.Context) {
	// validate id - check if integer
	id, err := CheckID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter valide number id"})
		return
	}

	// validation
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update user
	user, err := UpdateUserService(id, userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "User Updated Successfully", "data": user})
}

// delete user
func DeleteUser(c *gin.Context) {
	// validate id - check if integer
	id, err := CheckID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter valide number id"})
		return
	}

	Err := DeleteUserService(id)
	if Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully"})
}

// get emails 
func GetEmails(c *gin.Context) {
	var filterEmail FilterEmail
	err := c.ShouldBindJSON(&filterEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get all emails
	users, err := GetEmailsService(filterEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": users})
}






