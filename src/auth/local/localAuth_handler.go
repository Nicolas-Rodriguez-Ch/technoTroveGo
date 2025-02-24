package localAuth

import (
	"net/http"
	"strings"
	"technoTroveServer/src/api/users"
	"technoTroveServer/src/db"
	"technoTroveServer/src/models"
	"technoTroveServer/src/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func signUpHandler(c *gin.Context) {
	var input models.User
	var contactInfoArray []string
	var existingUser models.User

	input.Email = c.PostForm("email")

	err := db.DB.Where("email = ?", input.Email).First(&existingUser).Error

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "User with this email already exists"})
		return
	}

	input.Password = c.PostForm("password")
	input.FullName = c.PostForm("fullName")
	input.Description = c.PostForm("description")

	contactInfo := c.PostForm("contactInfo")
	if contactInfo != "" {
		contactInfoArray = strings.Split(contactInfo, ", ")
	}

	if input.Email == "" || input.Password == "" || input.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": "Invalid input"})
		return
	}

	if len(input.ContactInfo) > 0 && len(input.ContactInfo[0]) > 0 {
		contactInfoArray = strings.Split(input.ContactInfo[0], ",")

		for i, info := range contactInfoArray {
			contactInfoArray[i] = strings.TrimSpace(info)
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error hashing password", "error": err.Error()})
		return
	}

	profilePictures := utils.ConvertFilesToImageUrls(c)

	var profilePicture *string

	if len(profilePictures) > 0 {
		profilePicture = &profilePictures[0]
	}

	input.Password = string(passwordHash)
	input.ContactInfo = contactInfoArray
	input.ProfilePicture = profilePicture

	createdUser, err := users.CreateUser(&input, db.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user", "error": err.Error()})
		return
	}

	token, err := SignToken(&DecodedToken{ID: string(createdUser.ID)})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating token", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": gin.H{
			"fullName": createdUser.FullName,
			"email":    createdUser.Email,
		},
		"token": token,
	})
}
