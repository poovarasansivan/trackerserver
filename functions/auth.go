package functions

import (
	
	"net/http"
)

var Auth bool

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	// Perform header check/authentication here
	authHeader := r.Header.Get("Authorization")
	if authHeader != "YOUR_AUTH_TOKEN" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		Auth = false
		return
	}

	Auth = true
	// Call the next handler
	// http.DefaultServeMux.ServeHTTP(w, r)
}
