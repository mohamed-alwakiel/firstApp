package userModel

type User struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"user_name"`
	Email string `json:"email"`
	Age   uint   `json:"user_age"`
}

type UserInput struct {
	Name  string `json:"user_name" binding:"required,min=3,max=50"`
	Email string `json:"email" binding:"required,email"`
	Age   uint   `json:"user_age" binding:"min=18"`
}
