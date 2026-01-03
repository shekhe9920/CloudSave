package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

// Auth is middleware that checks if the request contains a valid JWT token.
// It verifies the user's authentication before allowing access to protected routes.
// Example: User closes the tab and comes back after 10 minutes,
// they will still be logged in as long as the token is valid
func Auth(secret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the Authorization header from the request
		auth := r.Header.Get("Authorization")

		// If no Authorization header is found, return Unauthorized error
		if auth == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Remove the "Bearer " part of the token string
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		// Parse the token and check its validity using the secret key
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		// If the token is invalid or expired, return Unauthorized error
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid, continue with the next handler
		next.ServeHTTP(w, r)
	})
}
