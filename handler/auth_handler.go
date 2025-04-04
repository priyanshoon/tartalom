package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"tartalom/config"
	"tartalom/database"
	"tartalom/model"
	"tartalom/utils"

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
	state := utils.GenerateState()
	url := config.GoogleOauthConfig().AuthCodeURL(state)
	return c.Status(http.StatusTemporaryRedirect).Redirect(url)
}

func LoginWithGooleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	state := c.Query("state")

	if code == "" || state == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

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

	// FIX: Fix the existing user issue: if login then server should generate token and send token as in response.
	var userExist model.User
	userFound := db.Where("google_id = ?", userInfo.ID).First(&userExist)

	if userFound.RowsAffected == 1 {
		token, err := utils.GenerateJWT(userExist)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal server error: error generating tokens")
		}

		response := fiber.Map{
			"token": token,
		}

		return c.Status(200).JSON(response)
	}

	password, err := utils.PasswordGenerator()
	if err != nil {
		log.Println("The password failed to generate")
	}

	user := model.User{
		ID:         uuid.New(),
		GoogleID:   userInfo.ID,
		Name:       userInfo.Name,
		Email:      userInfo.Email,
		Password:   password,
		Role:       "User",
		ProfilePic: userInfo.Picture,
	}

	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error inserting user into the database: %v", err)
		return c.Status(http.StatusInternalServerError).SendString("Error inserting user into the database.")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error: error generating tokens")
	}

	response := fiber.Map{
		"token": token,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func RegisterWithPassword(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Error fetching data: " + err.Error(),
		})
	}

	db := database.DB

	// TODO: Check user exist or not
	userExist := db.Where("email = ?", user.Email).First(&model.User{})

	if userExist.RowsAffected == 1 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already exist, go to login!",
		})
	}

	createUser := model.User{
		ID:       uuid.New(),
		GoogleID: "",
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     "User",
	}

	if err := db.Create(&createUser).Error; err != nil {
		log.Printf("Error inserting user into the database: %v", err)
		return c.Status(http.StatusInternalServerError).SendString("Error inserting user into the database.")
	}

	return c.SendString("Login with password")
}
