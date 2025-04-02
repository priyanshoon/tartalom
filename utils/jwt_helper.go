package utils

import (
	"log"
	"time"

	"tartalom/config"
	"tartalom/model"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user model.User) (string, error) {
	jwtSecret := config.GetJWTSecret()

	if jwtSecret == "" {
		log.Println("JWT Secret not set.")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
func ValidateJWT() {}
