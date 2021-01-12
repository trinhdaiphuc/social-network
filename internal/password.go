package internal

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

// HashPassword hash a password
func HashPassword(password string) string {
	cost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		cost = 7
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes)
}

// CheckPasswordHash check and hashed password correctly
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
