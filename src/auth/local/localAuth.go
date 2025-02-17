package localAuth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	localAuthGroup := router.Group("/auth/local")
	{
		localAuthGroup.POST("/signup" /*middleware.ProcessFileUploads, */, signUpHandler)
		// localAuthGroup.POST("/login", loginHandler)
	}
}
