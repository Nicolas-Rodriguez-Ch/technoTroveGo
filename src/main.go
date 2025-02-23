package main

import (
	"fmt"
	"os"
	"technoTroveServer/src/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dsn := os.Getenv("DATABASE_URL")
	
	db.Connect(dsn)

	println(port)
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	configeRoutes(router)

	router.Run(":" + port)
	fmt.Printf("Server is running on port %s", port)
}
