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

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDMyMjcwMjUsInVzZXJfaWQiOiIwMzk0ZmFkOC1hZTlhLTQxOWMtOGYyOC0wNmNkY2ViZWIwNTIifQ.I8r9sIrzGx4UCVf0c-bH0zWMyfXlBep5JBTObr88yWM
