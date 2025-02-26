package users

import (
	"technoTroveServer/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/", getAllUsersHandler)
		userGroup.GET("/profile", middleware.Auth, getUserByTokenHandler)
		userGroup.GET("/:id", getUserProfileHanlder)
		userGroup.PUT("", middleware.Auth, middleware.ProcessFileUpload, updateUserHandler)
		// userGroup.DELETE("/deactivate", middleware.Auth, deleteUser)
	}
}
