package main

import (
	"technoTroveServer/src/api/healthcheck"
	"technoTroveServer/src/api/projects"

	"github.com/gin-gonic/gin"
)

func configeRoutes(router *gin.Engine) {
	router.GET("/api/healthcheck", healthcheck.Handler)
	projects.RegisterRoutes(router)
	router.GET("/api/projects")
	router.POST("/auth/local")
}
