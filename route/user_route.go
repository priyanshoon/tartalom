package route

import (
	"tartalom/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func UserRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/", handler.Hello)
}
