package routes

import (
	"example/firstApp/routes/api"

	"github.com/gin-gonic/gin"
)

func MainRoute(r *gin.RouterGroup) {
	apiRoutes.UserRoute(r)
}
