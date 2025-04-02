package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

func PasswordGenerator() (string, error) {
	password := make([]byte, 20)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range password {
		index, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", fmt.Errorf("error generating index: %v", err)
		}
		password[i] = charset[index.Int64()]
	}

	return string(password), nil
}

func EncryptPassword() (string, error) {
	return "", nil
}
