package middleware

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func RequestId() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := uuid.New()
			fmt.Println(id.String())
			defer fmt.Println("after")
			h.ServeHTTP(w, r)
		})
	}
}
