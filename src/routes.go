package main

import (
	"technoTroveServer/src/api/healthcheck"
	"technoTroveServer/src/api/users"
	"technoTroveServer/src/auth/local"
	"technoTroveServer/src/api/projects"

	"github.com/gin-gonic/gin"
)

func configeRoutes(router *gin.Engine) {
	router.GET("/api/healthcheck", healthcheck.Handler)
	localAuth.RegisterRoutes(router)
	users.RegisterRoutes(router)
	projects.RegisterRoutes(router)
}
