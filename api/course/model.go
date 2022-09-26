package course

import (
	"example/firstApp/api/material"
	"example/firstApp/database"
)

type Course struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `gorm:"type:varchar(150);not null" json:"name"`
	Materials []material.Material
	//Users []*user.User `gorm:"many2many:user_courses;"`
}

// get all courses
func FindCourses() ([]Course, error) {
	courses := make([]Course, 0)
	if err := database.DB.Preload("Materials").Find(&courses).Error; err != nil {
		return courses, err
	}
	return courses, nil
}

// store new course
func Store(input Input) (Course, error) {
	course := Course{
		Name: input.Name,
	}

	if err := database.DB.Create(&course).Error; err != nil {
		return course, err
	}

	return course, nil
}

/*
// get specific user
func FindUser(id int) (User, error) {
	var user User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// update specific user
func Edit(user User, userInput UserInput) (User, error) {
	updatedUser := User{
		Name:  userInput.Name,
		Email: userInput.Email,
		Age:   userInput.Age,
	}

	// update user
	if err := database.DB.Model(&user).Updates(updatedUser).Error; err != nil {
		return user, err
	}

	return user, nil
}

// delete user
func Destroy(user User) error {

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

// get all emails
func FindEmails(filterEmail FilterEmail) ([]Email, error) {
	users := make([]Email, 0)

	query := EmailQuery(filterEmail)

	if err := database.DB.Model(User{}).Select("email").Where(query).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
*/
