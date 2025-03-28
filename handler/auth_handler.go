package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"tartalom/config"
	"tartalom/database"
	"tartalom/model"

	"github.com/gofiber/fiber/v2"
)

type GoogleUserInfo struct {
	ID       string
	Email    string
	Verified bool
	Name     string
	Picture  string
}

func LoginWithGoole(c *fiber.Ctx) error {
	url := config.GoogleConfig().AuthCodeURL("")
	return c.Status(http.StatusTemporaryRedirect).Redirect(url)
}

func LoginWithGooleCallback(c *fiber.Ctx) error {
	code := c.Query("code")

	tok, err := config.GoogleConfig().Exchange(c.Context(), code)
	if err != nil {
		log.Fatal(err)
	}

	client := config.GoogleConfig().Client(c.Context(), tok)

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

	user := model.User{
		ID:       userInfo.ID,
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Password: "",
		Role:     "User",
	}

	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error inserting user into the database: %v", err)
		return c.Status(http.StatusInternalServerError).SendString("Error inserting user into the database.")
	}

	return c.Status(http.StatusAccepted).JSON(userInfo)
}
