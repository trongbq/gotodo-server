package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/trongbq/gotodo-server/internal/api/auth"
)

func Auth(tokenVerifier *auth.TokenIssuer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := strings.Trim(r.Header.Get("Authorization"), " ")
			if token == "" {
				respondAuthFail(w, "Endpoint requires authentication token!")
				return
			}
			token = strings.TrimPrefix(token, "Bearer ")
			if token == "" {
				respondAuthFail(w, "Authentication token has invalid format!")
				return
			}
			claims, err := tokenVerifier.Verify(token)
			if err != nil {
				respondAuthFail(w, err.Error())
				return
			}
			ctx := context.WithValue(r.Context(), "UserID", claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func respondAuthFail(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(msg))
}
