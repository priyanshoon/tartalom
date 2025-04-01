package route

import (
	"tartalom/handler"
	"tartalom/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func BlogRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())

	blog := api.Group("/:user_id/blogs", middleware.AuthMiddleware())

	// get all blogs
	blog.Get("/", handler.GetBlogs)

	// get blog by blog_id

	blog.Post("/", handler.PostBlog)
	blog.Put("/", handler.UpdateBlog)
	blog.Delete("/", handler.DeleteBlog)
}
