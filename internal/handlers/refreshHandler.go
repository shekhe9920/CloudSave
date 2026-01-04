package handlers

import (
	"cloudsave/internal/config"
	"cloudsave/internal/middleware"
	"cloudsave/internal/utils"
	"encoding/json"
	"net/http"
)

// Refresh handles issuing a new access token using a valid refresh token.
//
// The refresh token is stored securely inside a server-side session,
// which is backed by an HttpOnly cookie. The client does not send the
// refresh token explicitly; it is retrieved from the session.
//
// Flow:
//  1. Retrieve the session from the global session store
//  2. Extract the refresh token from session values
//  3. Validate the refresh token
//  4. Extract the user ID (sub claim) from the token
//  5. Generate and return a new access token
func Refresh(w http.ResponseWriter, r *http.Request) {

	// Retrieve the session using the shared global CookieStore
	session, err := middleware.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Session error", http.StatusUnauthorized)
		return
	}

	// Extract the refresh token from the session
	refreshToken, ok := session.Values["refresh_token"].(string)
	if !ok || refreshToken == "" {
		http.Error(w, "No refresh token found", http.StatusUnauthorized)
		return
	}

	// Validate the refresh token and extract its claims
	claims, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	// Extract the subject (user ID) from claims
	subFloat, ok := claims["sub"].(float64)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	userID := int(subFloat)

	// Generate a new access token for the user
	cfg := config.Load()
	accessToken, err := utils.GenerateAccessToken(userID, cfg.JWTSecret)
	if err != nil {
		http.Error(w, "Failed to generate new access token", http.StatusInternalServerError)
		return
	}

	// Return the new access token to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
	})
}
