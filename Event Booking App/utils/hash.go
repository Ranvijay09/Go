package utils

import "golang.org/x/crypto/bcrypt"

func HashPaasword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPassword), err
}
