package middleware

import (
	"net/http"
)

type ApiKeyMiddleware struct {
	apiKey string
}

func NewApiKeyMiddleware(apiKey string) *ApiKeyMiddleware {
	return &ApiKeyMiddleware{
		apiKey: apiKey,
	}
}

func (akm *ApiKeyMiddleware) CheckAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		apiKey := req.Header.Get("api-key")
		if apiKey == akm.apiKey {
			next.ServeHTTP(w, req)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
