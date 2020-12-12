package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func dummyHandler(w http.ResponseWriter, r *http.Request) {}

func TestApiKeyMiddleware(t *testing.T) {
	// Arrange
	expectedStatusCode := http.StatusOK
	apiKey := "some_api_key"
	apiKeyMw := NewApiKeyMiddleware(apiKey)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/0701234567", nil)
	req.Header.Set("api-key", apiKey)

	router := mux.NewRouter()
	router.HandleFunc("/{phoneNumber}", dummyHandler).Methods("GET")
	router.Use(apiKeyMw.CheckAPIKey)
	// Act
	router.ServeHTTP(res, req)
	// Assert
	actualStatusCode := res.Result().StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("Got unexpected status code: %d, expected: %d", actualStatusCode, expectedStatusCode)
	}
}

func TestFailApiKeyMiddleware(t *testing.T) {
	// Arrange
	expectedStatusCode := http.StatusForbidden
	apiKey := "some_api_key"
	apiKeyMw := NewApiKeyMiddleware("some_other_api_key")

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/0701234567", nil)
	req.Header.Set("api-key", apiKey)

	router := mux.NewRouter()
	router.HandleFunc("/{phoneNumber}", dummyHandler).Methods("GET")
	router.Use(apiKeyMw.CheckAPIKey)
	// Act
	router.ServeHTTP(res, req)
	// Assert§
	actualStatusCode := res.Result().StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("Got unexpected status code: %d, expected: %d", actualStatusCode, expectedStatusCode)
	}
}
