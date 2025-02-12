package projects

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Project struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Images      []string `gorm:"type:text[]"`
	Active      bool
	User        struct {
		FullName string
	} `gorm:"foreignKey:UserID"`
}
