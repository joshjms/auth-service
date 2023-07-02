package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshjms/auth-service/database"
	"github.com/joshjms/auth-service/models"
)

func CreateUser(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	password = HashPassword([]byte(password))

	db := database.DB

	user := models.User{
		Username: username,
		Password: password,
	}

	res := db.Create(&user)

	if res.Error != nil {
		if res.Error.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			c.SendString("User already exists!")
			return c.SendStatus(400)
		} else {
			c.SendString("Failed to create user!")
			return c.SendStatus(500)
		}
	}

	return c.SendString("User created!")
}
