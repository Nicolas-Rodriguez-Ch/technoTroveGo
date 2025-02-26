package projects

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	projectGroup := router.Group("/api/projects")
	{
		projectGroup.GET("/", getAllProjectsHandler)
		// projectGroup.GET("/:id", GetProjectByID)
		// projectGroup.POST("/", middleware.Auth, middleware.ProcessFileUploads, CreateProject)
		// projectGroup.PUT("/:id", middleware.Auth, middleware.ProcessFileUploads, UpdateProject)
		// projectGroup.PATCH("/:id", middleware.Auth, DeleteProject)
	}
}
