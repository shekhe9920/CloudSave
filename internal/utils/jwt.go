package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateTokens generates a pair of JWT tokens for an authenticated user.
//
// It returns:
//   - A short-lived access token used for authenticating API requests
//   - A long-lived refresh token used to obtain new access tokens
//
// The access token is intended to be sent with requests (e.g. Authorization header),
// while the refresh token is stored securely server-side (via an HttpOnly session cookie).
func GenerateTokens(userID int, secret string) (string, string, error) {

	// Create access token claims
	accessClaims := jwt.MapClaims{
		"sub": userID,                                  // Subject (user ID)
		"exp": time.Now().Add(15 * time.Minute).Unix(), // Expires in 15 minutes
	}

	// Sign access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	// Create refresh token claims
	refreshClaims := jwt.MapClaims{
		"sub": userID,                                    // Subject (user ID)
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(), // Expires in 7 days
	}

	// Sign refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}

// GenerateAccessToken generates a new short-lived access token for a user.
//
// This function is typically used by the refresh endpoint after validating
// a refresh token. The generated access token contains the user ID as the
// subject (sub claim) and has a limited lifetime.
func GenerateAccessToken(userID int, secret string) (string, error) {

	// Create access token claims
	accessClaims := jwt.MapClaims{
		"sub": userID,                                  // Subject (user ID)
		"exp": time.Now().Add(15 * time.Minute).Unix(), // Expires in 15 minutes
	}

	// Sign access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedAccessToken, nil
}
