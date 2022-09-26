package course

import "github.com/gin-gonic/gin"

func CourseRoute(r *gin.RouterGroup) {
	route := r.Group("/courses")
	{
		route.GET("/", GetCourses)
		route.POST("/", CreateCourse)
		//route.GET("/:id", GetUser)
		//route.PUT("/:id", UpdateUser)
		//route.DELETE("/:id", DeleteUser)
	}
}
