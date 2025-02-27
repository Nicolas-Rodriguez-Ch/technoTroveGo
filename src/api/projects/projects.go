package projects

import (
	"technoTroveServer/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	projectGroup := router.Group("/api/projects")
	{
		projectGroup.GET("/", getAllProjectsHandler)
		projectGroup.GET("/:id", getProjectByIdHanlder)
		projectGroup.POST("", middleware.Auth, middleware.ProcessFileUpload, createProjectHandler)
		// projectGroup.PUT("/:id", middleware.Auth, middleware.ProcessFileUploads, UpdateProject)
		// projectGroup.PATCH("/:id", middleware.Auth, DeleteProject)
	}
}
