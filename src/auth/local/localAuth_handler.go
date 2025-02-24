package localAuth

import (
	"fmt"
	"net/http"
	"strings"
	"technoTroveServer/src/api/users"
	"technoTroveServer/src/db"
	"technoTroveServer/src/models"
	"technoTroveServer/src/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

	token, err := utils.SignToken(&utils.DecodedToken{ID: string(createdUser.ID)})

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

func loginHandler(c *gin.Context) {
	var loginInfo loginUser

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	user, err := utils.Login(db.DB, loginInfo.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email or password are incorrect"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email or password are incorrect"})
		return
	}

	token, err := utils.SignToken(&utils.DecodedToken{ID: fmt.Sprint(user.ID)})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating token", "error": err.Error()})
		return
	}

	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"data": gin.H{
			"email":    user.Email,
			"fullName": user.FullName,
		},
		"token": token,
	})
}
