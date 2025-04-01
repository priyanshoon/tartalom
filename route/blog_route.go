package route

import (
	"tartalom/handler"
	"tartalom/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func BlogRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())

	blog := api.Group("/user/blog", middleware.AuthMiddleware())

	blog.Get("/", handler.GetBlogs)
	blog.Post("/", handler.PostBlog)
	blog.Put("/", handler.UpdateBlog)
	blog.Delete("/", handler.DeleteBlog)
}
