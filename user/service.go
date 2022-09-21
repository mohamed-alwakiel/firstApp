package user

import (
	// "fmt"
	// "reflect"
	"strconv"
)

// get all users
func GetUsers(filterInput FilterInput) ([]User, error) {
	switch {
	case filterInput.Name != "" && filterInput.Email != "":
		return FindMultiFilter(filterInput)
	case filterInput.Name != "":
		return FindFilter("name", filterInput.Name)
	case filterInput.Email != "":
		return FindFilter("email", filterInput.Email)
	}
	return FindAll()
}

// create new user
func CreateUser(userInput UserInput) (User, error) {
	return Store(userInput)
}

// get specific user
func GetUser(id int) (User, error) {
	return FindUser(id)
}

// update specific user
func UpdateUser(id int, userInput UserInput) (User, error) {
	user, err := FindUser(id)
	if err != nil {
		return user, err
	}
	return Edit(user, userInput)
}

// delete specific user
func DeleteUser(id int) error {
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
