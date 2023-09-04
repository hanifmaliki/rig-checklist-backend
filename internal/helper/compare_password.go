package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
