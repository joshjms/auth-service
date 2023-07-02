package routes

import "github.com/gofiber/fiber/v2"

func ProtectedRoute(c *fiber.Ctx) error {
	return c.SendString("You are logged in!")
}
