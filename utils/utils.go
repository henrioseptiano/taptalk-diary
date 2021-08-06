package utils

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
