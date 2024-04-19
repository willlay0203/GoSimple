package middleware

import (
	"fmt"
	"net/http"
)

type Logger struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it down to the handler.
// This method comes from the interface of http.Handler
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Print("log middleware placeholder")
	l.handler.ServeHTTP(w, r)
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
