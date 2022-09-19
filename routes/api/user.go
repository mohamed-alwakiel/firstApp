package apiRoutes

import (
	"example/firstApp/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	route := r.Group("/users")
	{
		route.GET("/", controller.GetUsers)
		route.POST("/", controller.CreateUser)
		route.GET("/:id", controller.GetUser)
		route.PUT("/:id", controller.UpdateUser)
		route.DELETE("/:id", controller.DeleteUser)
	}
}
