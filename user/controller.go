package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get all users
func GetAll(c *gin.Context) {

	users, err := GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// create new user
func Create(c *gin.Context) {
	// validation
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store user
	user, err := CreateUser(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// get user by id
func Show(c *gin.Context) {
	// validate id - check if integer
	id, err := CheckID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter valide number id"})
		return
	}

	user, err := GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// update user
func Update(c *gin.Context) {
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
	user, err := UpdateUser(id, userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "User Updated Successfully", "data": user})
}

// delete user
func Delete(c *gin.Context) {
	// validate id - check if integer
	id, err := CheckID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter valide number id"})
		return
	}

	Err := DeleteUser(id)
	if Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully"})
}
