package middleware

import "net/http"

type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	// Needs to run in reverse to run the adapters (middlewares) in correct order
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}

	return h
}
