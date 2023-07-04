package auth

import (
	"github.com/joshjms/auth-service/database"
	"github.com/joshjms/auth-service/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(username, password string) (string, string, error) {
	db := database.DB
	var user models.User

	res := db.First(&user, "username = ?", username)

	if res.Error != nil {
		return "", "", res.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", err
	}

	access_token, err := CreateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refresh_token, err := CreateRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
}
