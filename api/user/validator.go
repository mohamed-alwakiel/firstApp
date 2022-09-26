package user

type Email struct {
	Email string `json:"email"`
}

type UserInput struct {
	Name  string `json:"name" binding:"required,min=3,max=50"`
	Email string `json:"email" binding:"required,email"`
	Age   uint   `json:"age" binding:"min=18"`
}

type FilterInput struct {
	Name  string `json:"name" `
	Email string `json:"email"`
}

type FilterEmail struct {
	Age    int    `json:"age"`
	Symbol string `json:"symbol"`
}
