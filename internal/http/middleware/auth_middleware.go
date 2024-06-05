package middleware

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authentication logic here
		// token := r.Header.Get("Authorization")

		next.ServeHTTP(w, r)
	})
}
