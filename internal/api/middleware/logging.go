package middleware

import (
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
	"github.com/trongbq/gotodo-server/internal/api/request"
)

func Logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Generate a request id
			rID := ksuid.New().String()
			// Put it in the log entry then put log entry to request context
			log := logrus.WithField("request-id", rID)
			ctx := request.WithLog(r.Context(), log)
			// Record start and end time of request to calculate latency
			start := time.Now()
			next.ServeHTTP(w, r.WithContext(ctx))
			end := time.Now()
			logrus.WithFields(logrus.Fields{
				"requestId":   rID,
				"method":      r.Method,
				"request":     r.RequestURI,
				"requestTime": start.Format(time.RFC3339),
				"latency":     end.Sub(start),
			}).Info()
		})
	}
}
