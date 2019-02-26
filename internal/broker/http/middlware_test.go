package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSercure(t *testing.T) {
	tt := []struct {
		name   string
		header map[string]string
		code   int
	}{
		{
			name: "ok",
			header: map[string]string{
				"Authorization": "Bearer token",
			},
			code: http.StatusOK,
		},
		{
			name:   "missing",
			header: map[string]string{},
			code:   http.StatusUnauthorized,
		},
		{
			name: "wrong format",
			header: map[string]string{
				"Authorization": "token",
			},
			code: http.StatusUnauthorized,
		},
		{
			name: "wrong token",
			header: map[string]string{
				"Authorization": "Bearer wrong",
			},
			code: http.StatusUnauthorized,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			b := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "http://exap,le.com", nil)
			for k, v := range tc.header {
				r.Header.Set(k, v)
			}
			h := Secure(b, "token")
			h.ServeHTTP(w, r)

			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected: %d", w.Code, tc.code)
			}
		})
	}
}
