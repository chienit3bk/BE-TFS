package middleware

import (
	"net/http"
	"project/packages/handlers/response"
)

const JsonContentType = "application/json"

func ContentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		if reqContentType != JsonContentType {
			response.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Request only allow content type: application/json !"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
