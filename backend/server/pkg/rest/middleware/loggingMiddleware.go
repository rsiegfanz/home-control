package middleware

import (
	"net/http"

	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.Logger.Info("New request", zap.String("RequestURI", r.RequestURI))
		next.ServeHTTP(w, r)
	})
}
