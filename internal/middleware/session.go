package middleware

import (
	"cloudsave/internal/config"
	"github.com/gorilla/sessions"
)

// Store is the global session store used across the application.
var Store *sessions.CookieStore

// InitSessionStore initializes the global CookieStore used for sessions.
//
// This function should be called exactly once during application startup
// (f.ex. in main.go).
func InitSessionStore() {
	cfg := config.Load()
	Store = sessions.NewCookieStore([]byte(cfg.JWTSecret))
}
