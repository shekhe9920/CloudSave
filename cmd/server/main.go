package main

import (
	"cloudsave/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"cloudsave/internal/config"
	"cloudsave/internal/db"
	"cloudsave/internal/middleware"
	"github.com/joho/godotenv"
)

func main() {

	// Initialize global CookieStore (DI)
	middleware.InitSessionStore()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load configuration from environment variables
	cfg := config.Load()

	// Print the loaded DB_PASS for debugging
	log.Println("DB Password:", cfg.DBPass) // Check that DB_PASS is loaded correctly
	fmt.Println("Connecting to database")

	// Connect to the database
	database, err := db.Connect(
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)
	if err != nil {
		log.Fatal("Database connection failed: %v", err)
	}
	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(w, r, database)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(w, r, database)
	})

	// New route for refreshing token
	mux.HandleFunc("/auth/refresh", handlers.Refresh)

	// New route for logging out
	mux.HandleFunc("/auth/logout", handlers.Logout)

	// protect dashboard-endpoint
	protected := middleware.Auth(cfg.JWTSecret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is protected content
		w.Write([]byte("Protected content"))
	}))

	// adding protected route
	mux.Handle("/api/dashboard", protected)

	handler := middleware.CORS(mux)

	log.Printf("Server running on :%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
