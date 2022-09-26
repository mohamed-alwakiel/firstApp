package material

type Input struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	CourseID uint   `json:"course_id" binding:"required"`
}
