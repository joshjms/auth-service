package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshjms/auth-service/auth"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	access_token, refresh_token, err := auth.LoginUser(username, password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"access_token":  access_token,
		"refresh_token": refresh_token,
	})
}
