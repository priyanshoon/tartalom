package handler

import (
	"log"

	"tartalom/model"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*model.User)
	if !ok {
		log.Println("gand mara looo")
	}
	log.Println(user.Email)
	return c.SendString("Api")
}
