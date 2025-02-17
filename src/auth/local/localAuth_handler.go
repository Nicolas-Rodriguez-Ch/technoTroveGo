package localAuth

import (
	"strings"
	"technoTroveServer/src/models"

	"github.com/gin-gonic/gin"
)

func signUpHandler(c *gin.Context) {
	var input models.User
	var contactInfoArray []string

	err := c.ShouldBind(&input)

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid input", "error": err})
	}

	if len(input.ContactInfo) > 0 && len(input.ContactInfo[0]) > 0 {
		contactInfoArray = strings.Split(input.ContactInfo[0], ",")

		for i, info := range contactInfoArray {
			contactInfoArray[i] = strings.TrimSpace(info)
		}
	}
}
