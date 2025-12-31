package main

import (
	"fmt"
	"log"
	"net/http"

	"cloudsave/internal/config"
	"cloudsave/internal/db"
)

func main() {
	cfg := config.Load()

	database, err := db.Connect(
		cfg.DBHost,
		cfg.DBUser,
		"CLOUDSAVE#99",
		cfg.DBName,
	)

	if err != nil {
		log.Fatal("Database connection failed: %v", err)
	}

	defer database.Close()

	fmt.Println("Connecting to database")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Server running on :%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
