package src

import "github.com/gin-gonic/gin"

func ConfigeRoutes(router *gin.Engine) {
	router.GET("api/healthCheck")
	router.GET("/api/users")
	router.GET("/api/projects")
	router.POST("/auth/local")
}
