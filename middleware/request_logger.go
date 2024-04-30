package middleware

import (
	"context"
	"log/slog"
	"net/http"
)

func LogRequest() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Context().Value(REQUESTID_CONTEXT_KEY)
			slog.Info("Request has been made",
				"requestId", requestId,
				"method", r.Method,
				"path", r.URL.Path,
			)
			ctx := context.WithValue(r.Context(), REQUESTID_CONTEXT_KEY, requestId)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
