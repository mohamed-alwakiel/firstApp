package material

import (
	"example/firstApp/database"
)

type Material struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `gorm:"type:varchar(150);not null" json:"name"`
	CourseID uint   `gorm:"" json:"course_id"`
	//Course   course.Course
}

// get all materials
func FindMaterials() ([]Material, error) {
	materials := make([]Material, 0)
	if err := database.DB.Find(&materials).Error; err != nil {
		return materials, err
	}
	return materials, nil
}

// store new material
func Store(input Input) (Material, error) {
	material := Material{
		Name:     input.Name,
		CourseID: input.CourseID,
	}

	if err := database.DB.Create(&material).Error; err != nil {
		return material, err
	}

	return material, nil
}
