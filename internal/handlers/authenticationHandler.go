package handlers

import (
	"cloudsave/internal/config"
	"cloudsave/internal/db"
	"cloudsave/internal/middleware"
	"cloudsave/internal/utils"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Login authenticates a user using email and password.
//
// On successful authentication:
//   - Generates a short-lived access token (JWT)
//   - Stores a long-lived refresh token in an HttpOnly session cookie
//   - Returns the access token in the JSON response body
//
// The refresh token is never exposed to JavaScript and is later used
// by the /auth/refresh endpoint to issue new access tokens.
func Login(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// All responses from this handler are JSON
	w.Header().Set("Content-Type", "application/json")

	// Structure for decoding login credentials from request body
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch user by email from database
	user, err := db.GetUserByEmail(database, credentials.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare stored password hash with provided password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(credentials.Password),
	); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate access and refresh tokens
	cfg := config.Load()
	accessToken, refreshToken, err := utils.GenerateTokens(user.ID, cfg.JWTSecret)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

	// Retrieve session from the global session store
	session, err := middleware.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	// Store refresh token in session
	session.Values["refresh_token"] = refreshToken

	// Configure session cookie options
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7, // 7 days
		HttpOnly: true,             // Not accessible via JavaScript
	}

	// Save session (this sets the Set-Cookie header)
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	// Return access token in response body
	json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
	})
}
