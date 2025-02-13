package projects

import "technoTroveServer/src/api/users"

// import (
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"log"
// )

type Project struct {
	ID          string       `gorm:"primaryKey"`
	Active      bool         `gorm:"default:true"`
	UserID      string       `gorm:"not null"`
	Title       string
	Description string       `gorm:"type:varchar(500)"`
	Images      []string     `gorm:"type:text[]"`
	Links       []string     `gorm:"type:text[]"`
	CreatedAt   string       `gorm:"default:current_timestamp"`
	UpdatedAt   string       `gorm:"default:current_timestamp"`
	User        users.User   `gorm:"foreignKey:UserID;references:ID"`
}