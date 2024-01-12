package utils

import (
	"crypto/sha256"
	"fmt"
)

func EncodeBase64(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashedInput := hash.Sum(nil)
	output := fmt.Sprintf("%x", hashedInput)
	return output
}

func HashSha256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashedInput := hash.Sum(nil)
	output := fmt.Sprintf("%x", hashedInput)
	return output
}
