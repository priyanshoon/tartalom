package utils

import (
	"fmt"
	"log"
	"time"

	"tartalom/config"
	"tartalom/model"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

func GenerateJWT(user model.User) (string, error) {
	jwtSecret := config.GetJWTSecret()

	if jwtSecret == "" {
		log.Println("JWT Secret not set.")
	}

	claims := &Claims{
		UserID: user.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	jwtSecret := config.GetJWTSecret()

	if jwtSecret == "" {
		log.Println("JWT Secret not set.")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}
