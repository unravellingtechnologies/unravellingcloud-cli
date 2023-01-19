package health

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestReadiness tests the readiness endpoint.
func TestReadiness(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/health/readyz", readiness)

	r.ServeHTTP(w, httptest.NewRequest("GET", "/health/readyz", nil))

	assert.Equal(t, w.Code, http.StatusOK, "Status code should be 200")

	var response ReadinessResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err, "Got an invalid JSON response", err, response)

	assert.True(t, assert.ObjectsAreEqualValues(response, ReadinessResponse{Ready: true, Message: "OK"}), "Response should be {\"ready\": true, \"message\": \"OK\"}")
}

// TestReadiness tests the readiness endpoint.
func TestHealthiness(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/health", healthiness)

	r.ServeHTTP(w, httptest.NewRequest("GET", "/health", nil))

	assert.Equal(t, w.Code, http.StatusOK, "Status code should be 200")

	var response HealthinessResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err, "Got an invalid JSON response", err, response)

	assert.True(t, assert.ObjectsAreEqualValues(response, HealthinessResponse{Healthy: true, Message: "OK"}), "Response should be {\"healthy\": true, \"message\": \"OK\"}")
}
