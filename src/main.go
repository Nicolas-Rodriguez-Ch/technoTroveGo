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
	mode := os.Getenv("GIN_MODE")
	autoMigrate := os.Getenv("AUTO_MIGRATE")

	if mode == "" {
		mode = gin.DebugMode
	}
	
	gin.SetMode(mode)
	db.Connect(dsn, autoMigrate)

	println(port)
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	configeRoutes(router)

	router.Run(":" + port)
	fmt.Printf("Server is running on port %s", port)
}
