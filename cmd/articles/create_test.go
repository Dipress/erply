package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Log("with initialized server and database.")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		s := setupServer("127.0.0.1:8080", "token", db)
		go s.ListenAndServe()
		defer s.Close()

		t.Log("\ttest:0\tshould create user.")
		{
			req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/articles", s.Addr), strings.NewReader(`{"title": "Title", "body": "Body"}`))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			req.Header.Set("Authorization", "Bearer token")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("unexpected status code: %d expected: %d", resp.StatusCode, http.StatusOK)
			}
		}
	}
}
