package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireRequestIDAllowsEnrichedRequest(t *testing.T) {
	calls := 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
	})

	wrapped := RequestID(RequireRequestID(handler))

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Request-ID", "req-12345")
	rr := httptest.NewRecorder()
	wrapped.ServeHTTP(rr, req)
	if calls != 1 {
		t.Errorf("Expected 1 call, got %d", calls)
	}
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestRequireRequestIDRejectsMissingID(t *testing.T) {
	calls := 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
	})

	wrapped := RequestID(RequireRequestID(handler))

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	wrapped.ServeHTTP(rr, req)
	if calls != 0 {
		t.Errorf("Expected 0 calls, got %d", calls)
	}
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}
