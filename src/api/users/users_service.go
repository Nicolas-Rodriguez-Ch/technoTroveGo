package users

type User struct {
	ID             uint `gorm:"primaryKey"`
	FullName       string
	Email          string
	Password       string
	Description    string
	ContactInfo    []string `gorm:"type:text[]"`
	ProfilePicture *string
}
