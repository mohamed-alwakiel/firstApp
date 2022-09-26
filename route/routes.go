package routes

import (
	"example/firstApp/api/course"
	"example/firstApp/api/material"
	"example/firstApp/api/user"

	"github.com/gin-gonic/gin"
)

func MainRoute(r *gin.RouterGroup) {
	user.UserRoute(r)
	course.CourseRoute(r)
	material.MaterialRoute(r)
}
