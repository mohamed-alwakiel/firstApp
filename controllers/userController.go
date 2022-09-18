package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-alwakiel/firstApp/models"
	"net/http"
)

// get all users
func GetUsers(c *gin.Context){
	var users=[]models.User
	
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H("data":users))
}