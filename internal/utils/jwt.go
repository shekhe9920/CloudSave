package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a JWT token for a user
// The token contains the user ID and expires after 24 hours
func GenerateJWT(userID int, secret string) (string, error) {

	// Data stored inside the token
	claims := jwt.MapClaims{
		"user_id": userID,                                // identifies the user
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // token expiration time
	}

	// Create a new JWT token using HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key and return it
	return token.SignedString([]byte(secret))
}
