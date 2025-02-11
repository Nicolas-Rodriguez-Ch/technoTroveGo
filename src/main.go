package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
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
