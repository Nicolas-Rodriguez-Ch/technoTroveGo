package users

import (
	"net/http"
	"technoTroveServer/src/db"

	"github.com/gin-gonic/gin"
)

func getAllUsersHandler(c *gin.Context) {
	users, err := getAllUsers(db.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved succesfully",
		"data":    users,
	})
}
