package route

import (
	"tartalom/handler"
	"tartalom/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func BlogRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())

	blog := api.Group("/:service_id/blog", middleware.AuthMiddleware())

	// Get all blogs
	blog.Get("/", handler.GetBlogs)

	blog.Post("/", handler.PostBlog)
	blog.Put("/:blog_id", handler.UpdateBlog)
	blog.Delete("/:blog_id", handler.DeleteBlog)
}
