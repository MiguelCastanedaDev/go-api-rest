package handlers

import (
	"net/http"

	"github.com/api-rest-go/internal/services"
	"github.com/api-rest-go/internal/types"
	"github.com/gin-gonic/gin"
)

func HandleGetAllSchools(c *gin.Context) {
	schools := services.GetAllSchoolsService()

	if len(schools) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Whoops! No schools found.", "status": http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": schools, "status": http.StatusOK})
}

func HandleCreateSchools(c *gin.Context) {
	var payload []types.CreateSchoolRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(payload) == 0 {
		c.JSON(400, gin.H{"error": "empty array"})
		return
	}

	created, err := services.CreateSchoolsService(payload)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, created)
}

func HandleDeleteSchool(c *gin.Context) {
	schoolID := c.Param("uuid")

	err := services.DeleteSchoolService(schoolID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete school.", "status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "School deleted successfully.", "status": http.StatusNoContent})
}

func HandleGetSchool(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name") // filtro opcional

	school, err := services.GetSchoolByIDService(id, name)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, school)
}
