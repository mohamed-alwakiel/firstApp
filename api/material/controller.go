package material

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// get all Courses
func GetMaterials(c *gin.Context) {
	// get all materials
	materials, err := GetMaterialsService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": materials})
}

// create new material
func CreateMaterial(c *gin.Context) {
	// validation
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store material
	material, err := CreateMaterialService(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": material})
}
