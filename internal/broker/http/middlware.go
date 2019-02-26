package http

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Secure represents middleware with authentication.
func Secure(after http.Handler, token string) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		authHdr := r.Header.Get("Authorization")
		if authHdr == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tknStr, err := parseAuthHeader(authHdr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if tknStr != token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		after.ServeHTTP(w, r)
	}

	return http.HandlerFunc(h)
}

// parseAuthHeader parses an authorization header. Expected header is of
// the format `Bearer <token>`.
func parseAuthHeader(bearerStr string) (string, error) {
	split := strings.Split(bearerStr, " ")
	if len(split) != 2 || strings.ToLower(split[0]) != "bearer" {
		return "", errors.New("xxpected Authorization header format: Bearer <token>")
	}

	return split[1], nil
}
