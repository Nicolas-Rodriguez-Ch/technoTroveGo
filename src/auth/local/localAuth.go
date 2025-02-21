package localAuth

import (
	"technoTroveServer/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	localAuthGroup := router.Group("/auth/local")
	{
		localAuthGroup.POST("/signup", middleware.ProcessFileUpload,  signUpHandler)
		// localAuthGroup.POST("/login", loginHandler)
	}
}
