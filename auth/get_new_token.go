package auth

import (
	"github.com/joshjms/auth-service/database"
	"github.com/joshjms/auth-service/models"
)

func GetNewToken(id float64) (string, error) {
	db := database.DB
	var user models.User

	res := db.First(&user, id)
	if res.Error != nil {
		return "", res.Error
	}

	access_token, err := CreateAccessToken(user)
	if err != nil {
		return "", err
	}

	return access_token, nil
}
