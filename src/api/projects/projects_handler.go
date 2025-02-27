package projects

import (
	"fmt"
	"net/http"
	"strings"
	"technoTroveServer/src/db"
	"technoTroveServer/src/models"
	"technoTroveServer/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/lucsky/cuid"
)

func getAllProjectsHandler(c *gin.Context) {
	projects, err := getAllProjects(db.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Projects retireved successfully",
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

func createProjectHandler(c *gin.Context) {
	user, exist := c.Get("user")
	id := user.(string)

	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	var input models.Project

	input.ID = cuid.New()
	input.Title = c.PostForm("title")
	input.Description = c.PostForm("description")
	input.UserID = id
	input.Active = true

	if input.Title == "" || input.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   "Invalid input",
		})
		return
	}

	links := c.PostForm("links")
	if links != "" {
		input.Links = pq.StringArray(strings.Split(links, ","))
		for i := range input.Links {
			input.Links[i] = strings.TrimSpace(input.Links[i])
		}
	}
	input.Images = pq.StringArray(utils.ConvertFilesToImageUrls(c))

	fmt.Println("Este es el valor de input antes de ser pasado a la creación:", input)
	createdProject, err := createProject(&input, db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating Project",
			"errpr":   err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
		"data":    createdProject,
	})
}
