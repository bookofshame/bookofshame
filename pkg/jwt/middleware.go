package jwt

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func (j Jwt) SetContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := getTokenFromHeader(r)
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		payload, err := j.Parse(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), dataCtxKey, payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := GetDataFromContext(r.Context())
		if err != nil || payload.UserId == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func getTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("authorization")
	parts := strings.Split(authHeader, " ")

	if len(parts) < 2 {
		return "", fmt.Errorf("invalid authorization header")
	}

	return parts[1], nil
}
