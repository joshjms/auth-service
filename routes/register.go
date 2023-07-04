package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshjms/auth-service/auth"
)

func Register(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	err := auth.CreateUser(username, password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
	})
}
