package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateState() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("Failed to generate random state: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
