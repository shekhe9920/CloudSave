// Package handlers contains the HTTP handler functions for the CloudSave API.
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"cloudsave/internal/db"
)

// Register handles the registration of a new user.
//
// It expects a JSON payload in the request body containing "email" and "password".
// On success, it creates a new user record in the database and returns the
// created user object as JSON.
//
// HTTP Responses:
//   - 201 Created: User was successfully created and returned in the body.
//   - 400 Bad Request: Invalid JSON input or missing required fields.
//   - 500 Internal Server Error: Database connection issues or failed JSON encoding.
func Register(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// set the response header for JSON
	w.Header().Set("Content-Type", "application/json")

	// parse incoming JSON data
	var userData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// check if email and password are provided
	if userData.Email == "" || userData.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
	}

	// create the user in the database
	user, err := db.CreateUser(database, userData.Email, userData.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}

/*
// TODO: Feature implementation.

// ConfirmEmail: Confirms user email after registration (future)
func ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	// Confirmation logic
}

// ResendConfirmation: Resends confirmation email (future)
func ResendConfirmation(w http.ResponseWriter, r *http.Request) {
	// Resend confirmation logic
}
*/
