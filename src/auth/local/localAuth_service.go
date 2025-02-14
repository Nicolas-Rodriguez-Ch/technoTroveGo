package localAuth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"errors"
	"os"
	"technoTroveServer/src/models"
	"time"
)

type DecodedToken struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func init() {
	godotenv.Load()
}

func Login(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SignToken(payload *DecodedToken) (string, error) {
	secret := os.Getenv("SECRET_KEY")

	if secret == "" {
		return "", errors.New("SECRET_KEY is not set in environment variables")
	}
	tokenExpirationTime := time.Now().Add(90 * 24 * time.Hour).Unix()

	claims := jwt.MapClaims{
		"id":  payload.ID,
		"exp": tokenExpirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (*DecodedToken, error) {
	secret := os.Getenv("SECRET_KEY")

	if secret == "" {
		return nil, errors.New("SECRET_KEY is not set in environment variables")
	}

	token, err := jwt.ParseWithClaims(tokenString, &DecodedToken{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*DecodedToken); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
