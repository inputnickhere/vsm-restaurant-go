package http

import (
	"net/http"
	"strings"
)

func StaticBearerAuth(token string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "server auth token is not configured", http.StatusInternalServerError)
				return
			}

			const prefix = "Bearer "
			if !strings.HasPrefix(h, prefix) {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			got := strings.TrimSpace(strings.TrimPrefix(h, prefix))
			if got != token {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
