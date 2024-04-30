package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const REQUESTID_CONTEXT_KEY ctxKey = "requestId"

func RequestId() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a new uuid and assign to the context
			requestId := uuid.New().String()
			ctx := context.WithValue(r.Context(), REQUESTID_CONTEXT_KEY, requestId)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
