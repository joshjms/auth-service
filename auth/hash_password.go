package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
