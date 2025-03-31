package handler

import (
	"time"

	"tartalom/database"
	"tartalom/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// FIX : Post Blog
func PostBlog(c *fiber.Ctx) error {
	// log.Println("Raw body:", string(c.Body()))
	blog := new(model.Blog)

	if err := c.BodyParser(blog); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Error fetching data: " + err.Error(),
		})
	}

	blog.Blog_ID = uuid.New()
	blog.PublishedDate = time.Now()

	db := database.DB

	db.Create(&blog)

	return c.Status(201).JSON(blog)
}

// TODO: Get All Blogs
func GetBlogs(c *fiber.Ctx) error {
	db := database.DB

	var blogs []model.Blog
	db.Find(&blogs)

	return c.Status(200).JSON(fiber.Map{
		"data": blogs,
	})
}

// TODO: Delete Blogs
func DeleteBlog(c *fiber.Ctx) error {
	blog := new(model.Blog)

	if err := c.BodyParser(blog); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Error fetching data: " + err.Error(),
		})
	}

	db := database.DB

	db.Delete(&model.Blog{}, blog.Blog_ID)

	return c.Status(200).JSON(fiber.Map{
		"message": "Blog Deleted Successfully",
	})
}

// TODO: Update Blogs
func UpdateBlog(c *fiber.Ctx) error {
	return c.Status(204).SendString("Updating Blogs...")
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDMyMzU4MTUsInVzZXJfaWQiOiI4MWZhYTUyNi02MmVkLTQyY2QtODBkMC1hYjQ2YWUyNzE4N2UifQ.5XH4WJKZW_XN624qjrH26C57XS2GeKBqUUskcEdn_N0
