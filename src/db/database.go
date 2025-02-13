package db

import (
	"log"
	"os"
	"technoTroveServer/src/api/projects"
	"technoTroveServer/src/api/users"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environentm")
	}

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	AutoMigrate()
}

func AutoMigrate() {
	err := DB.AutoMigrate(&users.User{}, &projects.Project{})
	if err != nil {
		log.Fatal("Error during auto migration:", err)
	}
	log.Println("Database auto-migration completed")
}
