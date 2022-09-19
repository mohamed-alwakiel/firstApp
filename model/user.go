package models

type User struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `gorm:"type:varchar(150);not null" json:"name"`
	Email string `gorm:"type:varchar(150);not null;unique" json:"email"`
	Age   uint   `json:"age"`
}
