package middleware

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// RequestId is a middleware handler that is ALWAYS on, it tags an ID to the request
type RequestId struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it down to the handler.
// This method comes from the interface of http.Handler
func (reqId *RequestId) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	fmt.Println(id.String())
	reqId.handler.ServeHTTP(w, r)
}

func NewRequestId(handlerToWrap http.Handler) *RequestId {
	return &RequestId{handlerToWrap}
}
