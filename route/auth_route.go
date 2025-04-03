package route

import (
	"tartalom/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	auth := api.Group("/auth")

	auth.Get("/login/google", handler.LoginWithGoole)
	auth.Get("/google/callback", handler.LoginWithGooleCallback)
	auth.Post("/register", handler.RegisterWithPassword)
}
