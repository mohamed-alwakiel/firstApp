package services

import (
	"example/firstApp/database"
	models "example/firstApp/model"
	"example/firstApp/validator"
	"strconv"
)

func GetUsers() ([]models.User, error) {
	users := make([]models.User, 0)

	if err := database.DB.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func CreateUser(userInput validator.UserInput) (models.User, error) {
	user := models.User{
		Name:  userInput.Name,
		Email: userInput.Email,
		Age:   userInput.Age,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUser(id int) (models.User, error) {
	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(id int, userInput validator.UserInput) (models.User, error) {
	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	updatedUser := models.User{
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
	var user models.User

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
