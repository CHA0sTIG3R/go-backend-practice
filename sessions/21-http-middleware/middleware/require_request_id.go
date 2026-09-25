package middleware

import "net/http"

func RequireRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID, ok := RequestIDFromContext(r.Context())
		if !ok || requestID == "" {
			http.Error(w, "Missing X-Request-ID header", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
