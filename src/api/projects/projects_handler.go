package projects

import (
	"net/http"
	"technoTroveServer/src/db"

	"github.com/gin-gonic/gin"
)

func getAllProjectsHandler(c *gin.Context) {
	projects, err := getAllProjects(db.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Projects retireved succesfully",
		"data":    projects,
	})
}
