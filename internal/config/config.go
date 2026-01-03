// Package config provides functions and types for managing application configuration.
// It handles loading settings from environment variables with sensible defaults.
package config

import "os"

// Config holds the application configuration settings.
type Config struct {
	Port      string // the network port the server listens on
	DBHost    string // database server address
	DBUser    string // database authentication username
	DBPass    string // database authentication password
	DBName    string // name of the database to connect to
	JWTSecret string // secret key used to sign JWT tokens
}

// Load initializes a new Config by fetching values from environment variables.
// If a variable is not set, it populates the field with a predefined default value.
// It returns a pointer to the populated Config struct.
func Load() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASS", ""),
		DBName:    getEnv("DB_NAME", "cloudsave"),
		JWTSecret: getEnv("JWT_SECRET", "dev-secret-cloud-save"),
	}
}

// getEnv searches for an environment variable with the name 'key'.
// If the variable exists, it returns the value, otherwise it returns the fallback.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
