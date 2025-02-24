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

func getUserByToken(c *gin.Context) {
	user, exist := c.Get("user")
	id := user.(string)
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	fetchedUser, err := getUserById(id, db.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User found!",
		"data":    fetchedUser,
	})
}
