package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type ctxKey string

func Adapt(h http.Handler, adapters ...Middleware) http.Handler {
	// Needs to run in reverse to run the adapters (middlewares) in correct order
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}

	return h
}
