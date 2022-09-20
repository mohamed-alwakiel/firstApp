package user

import (
	"example/firstApp/database"
	"strconv"
)

func GetUsers() ([]User, error) {
	users := make([]User, 0)

	if err := database.DB.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func CreateUser(userInput UserInput) (User, error) {
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

func GetUser(id int) (User, error) {
	var user User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(id int, userInput UserInput) (User, error) {
	var user User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

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

func DeleteUser(id int) error {
	var user User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func CheckID(id string) (int, error) {
	val, err := strconv.Atoi(id)

	if err != nil {
		return val, err
	}

	return val, nil
}
