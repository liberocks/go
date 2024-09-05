package middleware

import (
	"net/http"

	"github.com/liberocks/go/assignment/helpers"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get authorization from the request header
		authorization := r.Header.Get("Authorization")
		if authorization == "" || len(authorization) < 7 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Get the token from the authorization header (remove the "Bearer " prefix)
		accessToken := authorization[7:]

		// Validate the access token
		if err := helpers.VerifyAccessToken(accessToken); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
