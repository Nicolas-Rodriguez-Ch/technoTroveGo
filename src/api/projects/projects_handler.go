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

func getProjectByIdHanlder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No Project ID provided"})
		return
	}
	fetchedProject, err := getProjectById(id, db.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Project not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Project found!",
		"data":    fetchedProject,
	})
}
