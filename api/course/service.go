package course

// get all courses
func GetCoursesService() ([]Course, error) {
	return FindCourses()
}

// create new course
func CreateCourseService(input Input) (Course, error) {
	return Store(input)
}
