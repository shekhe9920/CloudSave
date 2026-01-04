package handlers

import (
	"cloudsave/internal/middleware"
	"net/http"
)

// Logout logs the user out by invalidating the session.
//
// This handler removes the refresh token by expiring the session cookie.
// After logout, the client will no longer be able to refresh access tokens
// and must authenticate again.
func Logout(w http.ResponseWriter, r *http.Request) {
	// Retrieve the existing session from the shared global session store
	session, err := middleware.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	// Expire the session cookie immediately
	session.Options.MaxAge = -1

	// Save session to apply cookie deletion
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to clear session", http.StatusInternalServerError)
		return
	}

	// Send confirmation response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}
