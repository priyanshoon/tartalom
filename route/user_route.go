package route

import (
	"tartalom/handler"
	"tartalom/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func UserRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())
	user := api.Group("/user", middleware.AuthMiddleware())

	user.Get("", handler.Hello)
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI2MzU0MDIsInVzZXJfaWQiOiJjNWZjY2FmYy1kNWNkLTQxZjYtYTBlMC04MTE5OTZhOTUxMmIifQ.piDXwE-db_d4Lb39MxYObyr-5SoRf5fsCzVKBStz8-I
