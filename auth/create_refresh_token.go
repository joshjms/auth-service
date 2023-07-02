package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joshjms/auth-service/models"
)

func CreateRefreshToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"type": "refresh",
		"uid":  user.ID,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret_key := os.Getenv("REFRESH_TOKEN_SIGNATURE")

	t, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}

	return t, err
}
