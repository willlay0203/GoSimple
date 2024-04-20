package middleware

import (
	"fmt"
	"net/http"
)

func LogRequest() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("request has been made")
			defer fmt.Println("after")
			h.ServeHTTP(w, r)
		})
	}
}
