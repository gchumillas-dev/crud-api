package handler

import (
	"net/http"
	"strings"
)

// AuthMiddleware verifies that the user was successful authorized.
func (env *Env) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := ""
		items := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(items) > 1 {
			token = items[1]
		}

		if len(token) == 0 {
			httpError(w, unauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// JSONMiddleware sets the content type to JSON.
func (env *Env) JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
