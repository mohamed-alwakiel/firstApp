package user

import (
	"reflect"
	"strconv"
)

// filter users
func UserQuery(filterInput FilterInput) string {
	return getFilter(filterInput)
}

func getAll() string {
	query := `name LIKE "%_%"`
	return query
}

func getFilter(filterInput FilterInput) string {

	values := reflect.ValueOf(filterInput)
	typesOf := values.Type()

	if values.NumField() <= 0 {
		return getAll()
	}

	var query string = ""

	for i := 0; i < values.NumField(); i++ {
		if i < values.NumField()-1 {
			query = query + typesOf.Field(i).Name + ` LIKE "%` + values.Field(i).String() + `%" AND `
			continue
		}
		query = query + typesOf.Field(i).Name + ` LIKE "%` + values.Field(i).String() + `%"`
	}

	return query
}

// filter emails
func EmailQuery(filterEmail FilterEmail) string {
	if filterEmail.Symbol == "smaller" {
		return getEmails("<", filterEmail.Age)
	}
	return getEmails(">", filterEmail.Age)
}

func getEmails(symbol string, value int) string {
	query := "age " + symbol + ` "` + strconv.Itoa(int(value)) + `"`
	return query
}
