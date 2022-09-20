package user

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	route := r.Group("/users")
	{
		route.GET("/", GetAll)
		route.POST("/", Create)
		route.GET("/:id", Show)
		route.PUT("/:id", Update)
		route.DELETE("/:id", Delete)
	}
}
