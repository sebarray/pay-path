package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a given password using bcrypt and returns the hashed password as a string.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
