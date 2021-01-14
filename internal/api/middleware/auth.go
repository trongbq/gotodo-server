package middleware

import (
	"context"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/trongbq/gotodo-server/internal/api/auth"
)

func Auth(tokenVerifier *auth.TokenIssuer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := log.WithFields(log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			})
			rtoken := strings.Trim(r.Header.Get("Authorization"), " ")
			if rtoken == "" {
				logger.Debug("auth token is empty")
				respondAuthFail(w, "Endpoint requires authentication token!")
				return
			}
			if !strings.HasPrefix(rtoken, "Bearer ") {
				logger.Debugf("auth token does not follow bearer format: %s", rtoken)
				respondAuthFail(w, "Authentication token has invalid format!")
				return
			}
			token := strings.TrimPrefix(rtoken, "Bearer ")
			claims, err := tokenVerifier.Verify(token)
			if err != nil {
				logger.Debugf("auth token can not be verified, error %s", err.Error())
				respondAuthFail(w, "Authentication token is invalid")
				return
			}
			logger.Debugf("authentication success for user %v", claims.UserID)
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
