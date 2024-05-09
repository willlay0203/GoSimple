package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	middleware "github.com/willlay0203/gohttp/middleware"
)

// Creates a new handler -> wraps new handler with RequestID() -> new handler checks for context
func TestRequestIdMiddleware(t *testing.T) {

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the request ID from the context
		ctx := r.Context()
		requestID, ok := ctx.Value(middleware.REQUESTID_CONTEXT_KEY).(string)
		if !ok {
			t.Error("Request ID not found in context")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if requestID == "" {
			t.Error("Empty request ID in context")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Respond with the request ID
		w.WriteHeader(http.StatusOK)
	})

	// Create a test request
	req := httptest.NewRequest("GET", "/", nil)

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Wrap the test handler with the middleware
	middleware.RequestId()(testHandler).ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
}
