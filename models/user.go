package userModel

type User struct {
	ID    uint    `json:"id" gorm:"primary_key"`
	Name  string  `json:"user_name"`
	Email *string `json:"email"`
	Age   uint    `json:"user_age"`
}

type CreateUserInput struct {
	Name  string  `json:"user_name" binding:"required"`
	Email *string `json:"email" binding:"required"`
	Age   uint    `json:"user_age"`
}
