package handlers

import (
	"cloudsave/internal/db"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Login checks user email and password and logs the user in
func Login(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// Tell the client that we return JSON
	w.Header().Set("Content-Type", "application/json")

	// Structure for incoming login data
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Read and decode JSON from request body
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find user in database by email
	user, err := db.GetUserByEmail(database, credentials.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare hashed password with the password provided by the user
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(credentials.Password),
	); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Login successful
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
