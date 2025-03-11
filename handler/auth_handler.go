package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"tartalom/config"
	"tartalom/database"
	"tartalom/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type GoogleUserInfo struct {
	ID       string
	Email    string
	Verified bool
	Name     string
	Picture  string
}

func LoginWithGoole(c *fiber.Ctx) error {
	url := config.GoogleOauthConfig().AuthCodeURL("")
	return c.Status(http.StatusTemporaryRedirect).Redirect(url)
}

func LoginWithGooleCallback(c *fiber.Ctx) error {
	code := c.Query("code")

	tok, err := config.GoogleOauthConfig().Exchange(c.Context(), code)
	if err != nil {
		log.Print(err)
		return c.Status(500).SendString("Lund lele mera!!!")
	}

	client := config.GoogleOauthConfig().Client(c.Context(), tok)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		log.Printf("Error making request to Google API: %v", err)
		return c.Status(http.StatusInternalServerError).SendString("Error making request to google")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Google API returned error: %s", resp.Status)
		return c.Status(resp.StatusCode).SendString("Error getting user info.")
	}

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("Error decoding user info: %v", err)
		return c.Status(http.StatusInternalServerError).SendString("Error decoding user info.")
	}

	db := database.DB

	// TODO: Check User exist or not
	userExist := db.Where("google_id = ?", userInfo.ID).First(&model.User{})

	if userExist.RowsAffected == 1 {
		return c.Status(404).SendString("User Exist")
	}

	user := model.User{
		ID:         uuid.New(),
		GoogleID:   userInfo.ID,
		Name:       userInfo.Name,
		Email:      userInfo.Email,
		Password:   "",
		Role:       "User",
		ProfilePic: userInfo.Picture,
	}

	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error inserting user into the database: %v", err)
		return c.Status(http.StatusInternalServerError).SendString("Error inserting user into the database.")
	}

	return c.Status(http.StatusAccepted).JSON(userInfo)
}

func LoginWithPassword(c *fiber.Ctx) error {
	return c.SendString("Login with password")
}
