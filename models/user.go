package models

type User struct {
	ID    uint    `json:"id" gorm:"primary_key"`
	Name  string  `json:"user_name"`
	Email *string `json:"email"`
	Age   uint    `json:"user_age"`
}


