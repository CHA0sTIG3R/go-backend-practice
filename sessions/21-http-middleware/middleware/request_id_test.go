package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Build a downstream http.HandlerFunc

/*
	read RequestIDFromContext(r.Context())

if absent:

	fail test

if present:

	write exact ID to response
*/
func TestRequestIDMiddlewarePassesIDThroughContext(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read RequestIDFromContext(r.Context())
		id, ok := RequestIDFromContext(r.Context())

		// if absent:
		if !ok {
			t.Errorf("Expected request ID not found in context")
		}

		// if present:
		if ok {
			// write exact ID to response
			w.Write([]byte(id))
		}
	})

	expectedID := "req-12345"

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Request-ID", expectedID)
	rr := httptest.NewRecorder()

	// Wrap the handler with the RequestID middleware
	middleware := RequestID(handler)

	// Serve the request
	middleware.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Check the response body for the expected request ID
	if rr.Body.String() != expectedID {
		t.Errorf("Expected request ID %q, got %q", expectedID, rr.Body.String())
	}
}
