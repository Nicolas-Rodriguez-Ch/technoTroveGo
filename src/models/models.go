package models

import (
	"github.com/lib/pq"
	"time"
)

type Project struct {
	ID          string         `gorm:"primaryKey"`
	Active      bool           `gorm:"default:true"`
	UserID      string         `gorm:"not null;index"`
	Title       string         `gorm:"not null"`
	Description string         `gorm:"type:varchar(500)"`
	Images      pq.StringArray `gorm:"type:text[]"`
	Links       pq.StringArray `gorm:"type:text[]"`
	CreatedAt   time.Time      `gorm:"type:timestamptz;default:now()"`
	UpdatedAt   time.Time      `gorm:"type:timestamptz;default:now()"`
	User        User           `gorm:"foreignKey:UserID;references:ID"`
}

type User struct {
	ID             string         `gorm:"primaryKey"`
	FullName       string         `gorm:"not null"`
	Email          string         `gorm:"unique;not null"`
	Password       string         `gorm:"not null"`
	Description    string         `gorm:"type:varchar(500)"`
	ContactInfo    pq.StringArray `gorm:"type:text[]"`
	ProfilePicture *string
	Auth0UserID    *string
	CreatedAt      time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt      time.Time `gorm:"type:timestamptz;default:now()"`
	Projects       []Project `gorm:"foreignKey:UserID;references:ID"`
}

type UserResponse struct {
	ID             string         `gorm:"primaryKey"`
	FullName       string         `json:"fullName"`
	Email          string         `gorm:"unique"`
	Description    string         `gorm:"type:varchar(500)"`
	ContactInfo    pq.StringArray `gorm:"type:text[]"`
	ProfilePicture *string
	Projects       []Project `gorm:"foreignKey:UserID;references:ID"`
}
