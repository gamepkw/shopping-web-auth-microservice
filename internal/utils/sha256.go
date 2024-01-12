package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func hashSHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func HashPassword(password string) string {
	secret := "shopping-web"

	passwordWithSalt := password + secret

	hashedPassword := hashSHA256(passwordWithSalt)
	return hashedPassword
}
