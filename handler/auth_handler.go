package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"tartalom/config"
	"tartalom/database"
	"tartalom/model"
	"tartalom/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
		jwtSecret := config.GetJWTSecret()
		if jwtSecret == "" {
			log.Println("JWT Secret not set.")
			return c.Status(fiber.StatusInternalServerError).SendString("Server configuration error.")
		}

		claims := jwt.MapClaims{
			"user_id": userExist.ID.String(),
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			log.Printf("error signing token : %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal server error: error generating tokens")
		}

		response := fiber.Map{
			"token": signedToken,
		}

		return c.Status(405).JSON(response)
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

	// generate jwt token
	jwtSecret := config.GetJWTSecret()
	if jwtSecret == "" {
		log.Println("JWT Secret not set.")
		return c.Status(fiber.StatusInternalServerError).SendString("Server configuration error.")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Printf("error signing token : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error: error generating tokens")
	}

	response := fiber.Map{
		"token": signedToken,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// func RegisterWithPassword(c *fiber.Ctx) error {
// 	user := new(model.User)
//
// 	if err := c.BodyParser(user); err != nil {
// 		return err
// 	}
//
// 	db := database.DB
//
// 	// TODO: Check user exist or not
// 	userExist := db.Where("email = ?", user.Email).First(&model.User{})
//
// 	if userExist.RowsAffected == 1 {
// 		return c.Status(302).SendString("User Exist")
// 	}
//
// 	postUser := model.User{
// 		ID:       uuid.New(),
// 		GoogleID: "",
// 		Name:     user.Name,
// 		Email:    user.Email,
// 		Password: user.Password,
// 		Role:     "User",
// 	}
//
// 	if err := db.Create(&postUser).Error; err != nil {
// 		log.Printf("Error inserting user into the database: %v", err)
// 		return c.Status(http.StatusInternalServerError).SendString("Error inserting user into the database.")
// 	}
//
// 	return c.SendString("Login with password")
// }
