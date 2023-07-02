package auth

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joshjms/auth-service/database"
	"github.com/joshjms/auth-service/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	db := database.DB
	var user models.User

	res := db.First(&user, "username = ?", username)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			log.Println(res.Error)
			c.SendString("User not found!")

			return c.SendStatus(400)
		} else {
			log.Println(res.Error)
			c.SendString("Failed to login!")

			return c.SendStatus(500)
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.SendString("Incorrect password!")
		return c.SendStatus(400)
	}

	access_token, err := CreateAccessToken(user)
	if err != nil {
		c.SendString("Failed to login!")

		return c.SendStatus(500)
	}

	refresh_token, err := CreateRefreshToken(user)
	if err != nil {
		c.SendString("Failed to login!")

		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{
		"access_token":  access_token,
		"refresh_token": refresh_token,
	})
}
