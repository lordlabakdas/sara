package handlers

import "github.com/gofiber/fiber/v2"


func GetUser(c *fiber.Ctx) error {
	return c.SendString("Get User")
}

func NewUser(c *fiber.Ctx) error {
	return c.SendString("New User")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
