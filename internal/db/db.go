// Package db provides primitives for database connectivity and management.
package db

import (
	"database/sql"
	"fmt"

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
