package route

import (
	"tartalom/handler"
	"tartalom/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func BlogRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())
	user := api.Group("/users/blog", middleware.AuthMiddleware())

	user.Get("", handler.Hello)
}
