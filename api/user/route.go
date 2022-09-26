package user

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	route := r.Group("/users")
	{
		route.GET("/", GetUsers)
		route.GET("/emails", GetEmails)
		route.POST("/", CreateUser)
		route.GET("/:id", GetUser)
		route.PUT("/:id", UpdateUser)
		route.DELETE("/:id", DeleteUser)
	}
}
