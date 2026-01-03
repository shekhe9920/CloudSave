package middleware

import "net/http"

// CORS allows the frontend (browser) to call this backend API.
// Without this, browser requests like fetch() will be blocked.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Allow requests from any website (for development)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow common HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Allow these headers in requests
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Browser sends OPTIONS request before some requests (preflight)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continue to the actual handler (login, register, etc.)
		next.ServeHTTP(w, r)
	})
}
