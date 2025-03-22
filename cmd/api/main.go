package main

import (
	"tartalom/database"
	"tartalom/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	route.AuthRoutes(app)
	route.UserRoute(app)

	app.Listen(":6969")
}
