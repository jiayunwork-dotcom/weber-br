package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func postJSON(t *testing.T, path string, body interface{}) *httptest.ResponseRecorder {
	t.Helper()
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewReader(data))
	rec := httptest.NewRecorder()
	Routes().ServeHTTP(rec, req)
	return rec
}

func TestWeberEndpoint(t *testing.T) {
	rec := postJSON(t, "/api/weber", map[string]interface{}{
		"rho_g": 1.2, "rho_l": 1000, "mu_l": 0.001, "sigma": 0.072, "U": 12, "d": 0.002,
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestUCritEndpoint(t *testing.T) {
	rec := postJSON(t, "/api/u-crit", map[string]interface{}{
		"rho_g": 1.2, "rho_l": 1000, "mu_l": 0.001, "sigma": 0.072, "d": 0.002,
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestInvalidReturns400(t *testing.T) {
	rec := postJSON(t, "/api/weber", map[string]interface{}{
		"rho_g": 1.2, "rho_l": 1000, "mu_l": 0.001, "sigma": 0.072, "U": -1, "d": 0.002,
	})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status=%d", rec.Code)
	}
}

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	Routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d", rec.Code)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/weber", nil)
	rec := httptest.NewRecorder()
	Routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status=%d", rec.Code)
	}
}
