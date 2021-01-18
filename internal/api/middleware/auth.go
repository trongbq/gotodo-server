package middleware

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/trongbq/gotodo-server/internal/api/auth"
	"github.com/trongbq/gotodo-server/internal/api/request"
	"github.com/trongbq/gotodo-server/internal/database"
)

func Auth(tokenVerifier *auth.TokenIssuer, db *database.DB) func(http.Handler) http.Handler {
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
				respondAuthFail(w, err.Error())
				return
			}
			logger.Debugf("authentication success for user %v", claims.UserID)

			// Authentication success, get current user and set to context
			user, err := db.GetUser(r.Context(), claims.UserID)
			if err != nil {
				if err == database.ErrNoRecordFound {
					respondAuthFail(w, "User not found")
					return
				}
				respondAuthFail(w, err.Error())
				return
			}

			next.ServeHTTP(w, r.WithContext(request.WithUser(r.Context(), user)))
		})
	}
}

func respondAuthFail(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(msg))
}
