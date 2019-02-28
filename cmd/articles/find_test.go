package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/romanyx/polluter"
)

const (
	findData = `
articles:
  - id: 1
    title: Title 
    body: Body
`
)

func TestFind(t *testing.T) {
	t.Log("with initialized server and database.")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		p := polluter.New(polluter.PostgresEngine(db))
		if err := p.Pollute(strings.NewReader(findData)); err != nil {
			t.Errorf("failed to pollute database: %v", err)
		}

		s := setupServer("127.0.0.1:8080", "token", db)
		go s.ListenAndServe()
		defer s.Close()

		t.Log("\ttest:0\tshould find article.")
		{
			req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/articles/1", s.Addr), nil)
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
