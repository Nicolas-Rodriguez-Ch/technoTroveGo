package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	println(port)
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	configeRoutes(router)

	
	router.Run(":" + port)
	fmt.Printf("Server is running on port %s", port)
}
