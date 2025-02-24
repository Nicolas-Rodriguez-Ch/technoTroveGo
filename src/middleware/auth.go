package middleware

import (
	"net/http"
	"strings"
	// "technoTroveServer/src/utils"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	if authorization == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "your session has expired, please log in again."})
		return
	}

	parts := strings.Split(authorization, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "your session has expired, please log in again."})
		return
	}

	// token := parts[1]
	// decodedToken, err := utils.VerifyToken(token)
}
