package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joshjms/auth-service/database"
	"github.com/joshjms/auth-service/models"
)

func GetNewToken(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	id := claims["uid"].(float64)

	db := database.DB
	var user models.User

	res := db.First(&user, id)
	if res.Error != nil {
		c.SendString("Failed to get user!")
		return c.SendStatus(500)
	}

	access_token, err := CreateAccessToken(user)
	if err != nil {
		c.SendString("Failed to get new token!")
		return c.SendStatus(500)
	}

	c.JSON(fiber.Map{
		"access_token": access_token,
	})

	return nil
}
