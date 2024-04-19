package server_test

import (
	"bytes"
	"errors"
	. "goroute/server"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const TEST_SUCCESS_STRING = "Test success"

// Unit test
func TestSetup(t *testing.T) {
	m := Setup(":8080")

	if m.Port != ":8080" {
		t.Fatal("Port value different")
	}
}

// Unit test happy cases
func TestRouteSetters(t *testing.T) {
	methods := [5]string{"get", "post", "put", "patch", "delete"}

	for _, m := range methods {
		// Create a recorder to capture the response
		recorder := httptest.NewRecorder()

		testServer := Setup(":8080")

		// Register the handlers with the test path
		testServer.GET("/test", mockHandler)
		testServer.POST("/test", mockHandler)
		testServer.PUT("/test", mockHandler)
		testServer.PATCH("/test", mockHandler)
		testServer.DELETE("/test", mockHandler)

		switch m {
		case "get":
			// Create a dummy request for the test path
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			// Serve the request
			testServer.Mux.ServeHTTP(recorder, req)
			log.Printf("Test completed for method %s", m)
		case "post":
			// Create a dummy request for the test path
			req, err := http.NewRequest("POST", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			// Serve the request
			testServer.Mux.ServeHTTP(recorder, req)
			log.Printf("Test completed for method %s", m)
		case "put":
			// Create a dummy request for the test path
			req, err := http.NewRequest("PUT", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			// Serve the request
			testServer.Mux.ServeHTTP(recorder, req)
			log.Printf("Test completed for method %s", m)
		case "patch":
			// Create a dummy request for the test path
			req, err := http.NewRequest("PATCH", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			// Serve the request
			testServer.Mux.ServeHTTP(recorder, req)
			log.Printf("Test completed for method %s", m)
		case "delete":
			// Create a dummy request for the test path
			req, err := http.NewRequest("DELETE", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			// Serve the request
			testServer.Mux.ServeHTTP(recorder, req)
			log.Printf("Test completed for method %s", m)
		}

		if err := validateResponse(recorder); err != nil {
			t.Fatal(err)
		}
	}
}

func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(TEST_SUCCESS_STRING))
}

func validateResponse(recorder *httptest.ResponseRecorder) (err error) {
	// Check the response code
	if status := recorder.Code; status != http.StatusOK {
		return errors.New("Handler returned NOT 200")
	}

	// Get recorder body as byte slice
	responseBody := recorder.Body.Bytes()

	// Check if the response body contains the expected string from the mock handler
	if !bytes.Contains(responseBody, []byte(TEST_SUCCESS_STRING)) {
		return errors.New("Handler did not return expected string")
	}

	return nil
}
