package user

import (
	"example/firstApp/database"
)

type User struct {
	ID    uint   `gorm:"primary_key" json:"id" filter:"searchable;filterable"`
	Name  string `gorm:"type:varchar(150);not null" json:"name" filter:"searchable;filterable"`
	Email string `gorm:"type:varchar(150);not null;unique" json:"email" filter:"searchable;filterable"`
	Age   uint   `json:"age"`
}

// get all users
func FindAll() ([]User, error) {
	users := make([]User, 0)
	if err := database.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// get filterd users with specific col
func FindFilter(field string, value string) ([]User, error) {
	users := make([]User, 0)
	if err := database.DB.Where(field+" LIKE ?", "%"+value+"%").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// get filterd user with multi col
func FindMultiFilter(filterInput FilterInput) ([]User, error) {
	users := make([]User, 0)
	if err := database.DB.Where("name LIKE ? AND email LIKE ?", "%"+filterInput.Name+"%", "%"+filterInput.Email+"%").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// store new user
func Store(userInput UserInput) (User, error) {
	user := User{
		Name:  userInput.Name,
		Email: userInput.Email,
		Age:   userInput.Age,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

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
