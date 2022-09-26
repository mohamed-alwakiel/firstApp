package course

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// get all Courses
func GetCourses(c *gin.Context) {
	// get all courses
	users, err := GetCoursesService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// create new course
func CreateCourse(c *gin.Context) {
	// validation
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store course
	course, err := CreateCourseService(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": course})
}
