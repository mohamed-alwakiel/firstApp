package material

import "github.com/gin-gonic/gin"

func MaterialRoute(r *gin.RouterGroup) {
	route := r.Group("/materials")
	{
		route.GET("/", GetMaterials)
		route.POST("/", CreateMaterial)
		//route.GET("/:id", GetUser)
		//route.PUT("/:id", UpdateUser)
		//route.DELETE("/:id", DeleteUser)
	}
}
