package projects

import (
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

func updateProjectHandler(c *gin.Context) {
	user, exist := c.Get("user")
	projectId := c.Param("id")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	if projectId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No Project ID provided"})
		return
	}
	id, ok := user.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID"})
		return
	}

	ownerShipStatus, err := checkProjectOwnerShip(id, projectId, db.DB)

	if ownerShipStatus != 200 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You don't have permission to update this project.",
			"error":   err,
		})
		return
	}

	var input models.Project
	input.Title = c.PostForm("title")
	input.Description = c.PostForm("description")

	links := c.PostForm("links")
	if links != "" {
		input.Links = pq.StringArray(strings.Split(links, ","))
		for i := range input.Links {
			input.Links[i] = strings.TrimSpace(input.Links[i])
		}
	}
	input.Images = pq.StringArray(utils.ConvertFilesToImageUrls(c))
	updatedProject, err := updateProject(id, &input, db.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error updating Project",
			"Error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Project updated",
		"data":    updatedProject,
	})
}
