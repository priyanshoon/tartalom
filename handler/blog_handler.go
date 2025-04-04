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
	user_id := c.Params("user_id")
	blog := new(model.Blog)

	if err := c.BodyParser(blog); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Error fetching data: " + err.Error(),
		})
	}

	blog.Blog_ID = uuid.New()
	blog.PublishedDate = time.Now()
	blog.UserID = uuid.MustParse(user_id)

	db := database.DB

	db.Create(&blog)

	return c.Status(201).JSON(blog)
}

// FIX: Get All Blogs
func GetBlogs(c *fiber.Ctx) error {
	db := database.DB

	var blogs []model.Blog
	db.Find(&blogs)

	return c.Status(200).JSON(blogs)
}

// FIX: Delete Blogs
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

// FIX: Update Blogs
func UpdateBlog(c *fiber.Ctx) error {
	blog := new(model.Blog)

	if err := c.BodyParser(blog); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Error fetching data: " + err.Error(),
		})
	}

	db := database.DB

	result := db.Model(&model.Blog{}).Where("blog_id = ?", blog.Blog_ID).Updates(&model.Blog{
		Title: blog.Title,
		Body:  blog.Body,
	})

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server error",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "blog not found!",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Blog Updated!",
	})
}

// TODO: Get blog by ID (blog_id)
func GetBlogById(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{})
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM1OTk1MTAsInVzZXJfaWQiOiI4MWZhYTUyNi02MmVkLTQyY2QtODBkMC1hYjQ2YWUyNzE4N2UifQ.0O6ZnaNsi3UBJSndoFVlAgUmq2kqTdzzqBx16uJ8AZc
