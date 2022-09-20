package routes

import (
	"example/firstApp/user"

	"github.com/gin-gonic/gin"
)

func MainRoute(r *gin.RouterGroup) {
	user.UserRoute(r)
}
