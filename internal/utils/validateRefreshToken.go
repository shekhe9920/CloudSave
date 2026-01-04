package utils

import (
	"cloudsave/internal/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

// ValidateRefreshToken validates a refresh token and returns its claims.
//
// The function verifies that:
//   - The token is correctly signed using the application's JWT secret
//   - The token is not expired
//   - The token is structurally valid
//
// On success, the decoded JWT claims are returned. On failure, an error
// is returned and the token should be considered invalid.
func ValidateRefreshToken(refreshToken string) (jwt.MapClaims, error) {
	// Load configuration to access the JWT signing secret
	cfg := config.Load()

	// Prepare a MapClaims object to store decoded token claims
	claims := jwt.MapClaims{}

	// Parse and validate the refresh token
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	// Return validated claims
	return claims, nil
}
