package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Get All Books")
}


func NewBook(c *fiber.Ctx) error {
	return c.SendString("New Book")
}

