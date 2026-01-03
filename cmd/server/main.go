package main

import (
	"cloudsave/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"cloudsave/internal/config"
	"cloudsave/internal/db"
	"cloudsave/internal/middleware"
)

func main() {
	// load inn configuration
	cfg := config.Load()

	database, err := db.Connect(
		cfg.DBHost,
		cfg.DBUser,
		"CLOUDSAVE#99",
		cfg.DBName,
	)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	defer database.Close()

	fmt.Println("Connecting to database")

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

	handler := middleware.CORS(mux)

	log.Printf("Server running on :%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
