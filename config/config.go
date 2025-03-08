package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env")
	}

	return os.Getenv(key)
}

func GoogleOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     Config("GOOGLE_CLIENT_ID"),
		ClientSecret: Config("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:6969/api/auth/google/callback",
		Scopes: []string{
			"email",
			"profile",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
}
