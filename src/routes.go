package main

import (
	"technoTroveServer/src/api/healthcheck"

	"github.com/gin-gonic/gin"
)

func configeRoutes(router *gin.Engine) {
	router.GET("/api/healthcheck", healthcheck.Handler)
	router.GET("/api/users")
	router.GET("/api/projects")
	router.POST("/auth/local")
}
