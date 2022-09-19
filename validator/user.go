package validator

import (
	// "net/http"

	// "github.com/gin-gonic/gin"
)

type UserInput struct {
	Name  string `json:"name" binding:"required,min=3,max=50"`
	Email string `json:"email" binding:"required,email"`
	Age   uint   `json:"age" binding:"min=18"`
}

