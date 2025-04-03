package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

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

	hashPassword, err := utils.EncryptPassword(password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	user := model.User{
		ID:         uuid.New(),
		GoogleID:   userInfo.ID,
		Name:       userInfo.Name,
		Email:      userInfo.Email,
		Password:   hashPassword,
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

	if user.Name == "" || len(user.Name) < 2 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Name should be greater than 2 char",
		})
	}

	validEmail, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error validating email",
		})
	}

	if !validEmail {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Email",
		})
	}

	if len(user.Password) < 4 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password should contain at least 4 characters",
		})
	}

	db := database.DB

	var userExist model.User
	userFound := db.Where("email = ?", user.Email).First(&userExist)

	if userExist.GoogleID != "" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Google account exist with the associating email id, please login with google.",
		})
	}

	if userFound.RowsAffected == 1 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already exist, go to login!",
		})
	}

	password, err := utils.EncryptPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	createUser := model.User{
		ID:       uuid.New(),
		GoogleID: "",
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
		Role:     "User",
	}

	if err := db.Create(&createUser).Error; err != nil {
		log.Printf("Error inserting user into the database: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user registered successfully",
	})
}

// TODO: create handler for manual login for users
func LoginWithPassword(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Error fetching data: " + err.Error(),
		})
	}

	validEmail, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error validating email",
		})
	}

	if !validEmail {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Email",
		})
	}

	db := database.DB

	var userExist model.User
	userFound := db.Where("email = ?", user.Email).First(&userExist)

	if userFound.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	if userExist.GoogleID != "" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Google account exist with the associating email id, please login with google.",
		})
	}

	validPass := utils.ValidateHashPassword(userExist.Password, user.Password)
	if !validPass {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid login credentials",
		})
	}

	token, err := utils.GenerateJWT(userExist)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error: error generating tokens")
	}

	response := fiber.Map{
		"token": token,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
