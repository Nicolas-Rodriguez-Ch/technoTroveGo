package healthcheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server running fine"})
}
