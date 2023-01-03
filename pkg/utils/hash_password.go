package utils

import "golang.org/x/crypto/bcrypt"

// Convert to Hash
func HashText(text string) string {
	newText, _ := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)
	return string(newText)
}

// Compare hash and password
func CompareHashPassword(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}
