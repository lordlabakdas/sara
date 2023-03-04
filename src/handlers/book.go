package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lordlabakdas/sara/src/db/service"
	"gorm.io/gorm"
)

func GetBooks(c *fiber.Ctx) error {
	service.GetBook(db *gorm.DB)
	return c.SendString("Get All Books")
}


func NewBook(c *fiber.Ctx) error {
	return c.SendString("New Book")
}

