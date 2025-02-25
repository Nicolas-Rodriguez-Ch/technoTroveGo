package users

import (
	"net/http"
	"strings"
	"technoTroveServer/src/db"
	"technoTroveServer/src/models"
	"technoTroveServer/src/utils"

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

func getUserByTokenHandler(c *gin.Context) {
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

func getUserProfileHanlder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No user ID provided"})
		return
	}
	fetchedUser, err := getUserProfile(id, db.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User found!",
		"data":    fetchedUser,
	})
}

func updateUserHandler(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	id, ok := user.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID"})
		return
	}

	var input models.User
	var contactInfoArray []string

	input.Password = c.PostForm("password")
	input.FullName = c.PostForm("fullName")
	input.Description = c.PostForm("description")

	contactInfo := c.PostForm("contactInfo")
	if contactInfo != "" {
		contactInfoArray = strings.Split(contactInfo, ", ")
		input.ContactInfo = contactInfoArray
	}

	profilePictures := utils.ConvertFilesToImageUrls(c)
	if len(profilePictures) > 0 {
		input.ProfilePicture = &profilePictures[0]
	}

	updatedUser, err := updateUser(id, &input, db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error updating user",
			"Error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated",
		"data":    updatedUser,
	})
}
