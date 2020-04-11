package middleware

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/globalconstant"
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO allow only from whitelisted servers in config
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Get auth token
		token := r.Header.Get(globalconstant.AuthHeader)
		ctx := context.WithValue(r.Context(), globalconstant.TokenKey, token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
