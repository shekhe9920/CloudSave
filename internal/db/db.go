// Package db provides primitives for database connectivity and management.
package db

import (
	"cloudsave/internal/models"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Connect establishes a new connection pool to MySQL database using the provided credentials.
//
// Constructs a DSN (Data Source Name) from the host, user, password,and database name.
// The configuration is configured with parseTime=true to automatically handle MySQL date/time
// types as Go 'time.Time' objects.
//
// This function performs a Ping to verify that the database server i reachable and
// that the credentials are valid before returning the *sql.DB instance.
//
// It returns an error if the driver fails to initialize or if the connection test fails.
func Connect(host, user, password, name string) (*sql.DB, error) {
	// format the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		user,
		password,
		host,
		name,
	)

	// validate the DSN format
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// verify the connection is actually alive by pinging the database server
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}

// CreateUser creates a new user in the database
func CreateUser(db *sql.DB, email, password string) (*models.User, error) {
	// hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// insert the user into the database
	result, err := db.Exec("INSERT INTO users (email, password_hash) VALUES (?, ?)", email, string(hashPassword))
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// get the user ID
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get user ID: %v", err)
	}

	// return the user
	user := &models.User{
		ID:           int(userID),
		Email:        email,
		PasswordHash: string(hashPassword),
		CreatedAt:    time.Now(),
	}

	return user, nil
}

// GetUserByEmail fetches a user from the database using the email address
func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {

	// SQL query to find a user by email
	query := `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = ?
	`

	// Run the query with the provided email
	row := db.QueryRow(query, email)

	// Create an empty User struct to store the result
	user := &models.User{}

	// Read the database row into the user struct
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		// No user found with this email
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		// Any other database error
		return nil, err
	}

	// User found successfully
	return user, nil
}
