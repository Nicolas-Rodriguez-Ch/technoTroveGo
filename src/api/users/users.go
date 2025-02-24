package users

import (
	"technoTroveServer/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/", getAllUsersHandler)
		userGroup.GET("/profile", middleware.Auth, getUserByToken)
		// userGroup.GET("/:id", getUserProfile)
		// userGroup.PUT("/", middleware.Auth, middleware.ProcessFileUploads, updateUser)
		// userGroup.DELETE("/deactivate", middleware.Auth, deleteUser)
	}
}
