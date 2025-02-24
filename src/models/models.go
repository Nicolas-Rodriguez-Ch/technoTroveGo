package models

import (
	"github.com/lib/pq"
	"time"
)

type Project struct {
	ID          uint `gorm:"primaryKey"`
	Active      bool `gorm:"default:true"`
	UserID      uint `gorm:"index"`
	Title       string
	Description string    `gorm:"type:varchar(500)"`
	Images      []string  `gorm:"type:text[]"`
	Links       []string  `gorm:"type:text[]"`
	CreatedAt   time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt   time.Time `gorm:"type:timestamptz;default:now()"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
}

type User struct {
	ID             uint `gorm:"primaryKey"`
	FullName       string
	Email          string `gorm:"unique"`
	Password       string
	Description    string         `gorm:"type:varchar(500)"`
	ContactInfo    pq.StringArray `gorm:"type:text[]"`
	ProfilePicture *string
	Auth0UserID    *string
	CreatedAt      time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt      time.Time `gorm:"type:timestamptz;default:now()"`
	Projects       []Project `gorm:"foreignKey:UserID;references:ID"`
}
