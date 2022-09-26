package user

import (
	// "fmt"
	// "reflect"
	"strconv"
)

// get all users
func GetUsersService(filterInput FilterInput) ([]User, error) {
	return FindUsers(filterInput)
}

// create new user
func CreateUserService(userInput UserInput) (User, error) {
	return Store(userInput)
}

// get specific user
func GetUserService(id int) (User, error) {
	return FindUser(id)
}

// update specific user
func UpdateUserService(id int, userInput UserInput) (User, error) {
	user, err := FindUser(id)
	if err != nil {
		return user, err
	}
	return Edit(user, userInput)
}

// delete specific user
func DeleteUserService(id int) error {
	user, err := FindUser(id)
	if err != nil {
		return err
	}
	return Destroy(user)
}

// check id validate - must be integer
func CheckID(id string) (int, error) {
	val, err := strconv.Atoi(id)

	if err != nil {
		return val, err
	}

	return val, nil
}

// get all users
func GetEmailsService(filterEmail FilterEmail) ([]Email, error) {
	return FindEmails(filterEmail)
}
