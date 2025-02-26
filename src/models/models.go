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
	ID             string         `json:"id"`
	FullName       string         `json:"fullName"`
	Email          string         `json:"email"`
	Description    string         `json:"description"`
	ProfilePicture *string        `json:"profilePicture"`
	ContactInfo    pq.StringArray `json:"contactInfo" gorm:"type:text[]"`
	Projects       []Project      `json:"projects" gorm:"foreignKey:UserID;references:ID"`
}

type ProjectResponse struct {
	ID          string         `json:"id"`
	Active      bool           `json:"active"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Links       pq.StringArray `json:"links" gorm:"type:text[]"`
	User        User           `json:"user"`
	UserID      string
}
