package handler

import "github.com/gofiber/fiber/v2"

// TODO: Post Blog
func PostBlog(c *fiber.Ctx) error {
	return c.Status(201).SendString("Posting Blogs...")
}

// TODO: Get Blogs
func GetBlogs(c *fiber.Ctx) error {
	return c.Status(200).SendString("Getting Blogs...")
}

// TODO: Delete Blogs
func DeleteBlog(c *fiber.Ctx) error {
	return c.Status(200).SendString("Deleting Blogs...")
}

// TODO: Update Blogs
func UpdateBlog(c *fiber.Ctx) error {
	return c.Status(204).SendString("Updating Blogs...")
}
