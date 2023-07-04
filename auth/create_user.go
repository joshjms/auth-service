package auth

import (
	"github.com/joshjms/auth-service/database"
	"github.com/joshjms/auth-service/models"
)

func CreateUser(username, password string) error {
	password = HashPassword([]byte(password))

	db := database.DB

	user := models.User{
		Username: username,
		Password: password,
	}

	res := db.Create(&user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
