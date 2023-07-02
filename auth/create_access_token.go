package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joshjms/auth-service/models"
)

func CreateAccessToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"uid":  user.ID,
		"name": user.Username,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret_key := os.Getenv("ACCESS_TOKEN_SIGNATURE")

	t, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}

	return t, err
}
