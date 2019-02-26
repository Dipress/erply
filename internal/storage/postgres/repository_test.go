package postgres

import (
	"context"
	"strings"
	"testing"

	"github.com/romanyx/erply/internal/article"
	"github.com/romanyx/polluter"
)

func TestRepositoryCreate(t *testing.T) {
	t.Parallel()
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		r := NewRepository(db)

		t.Log("\ttest:0\tshould insert model into database")
		{
			na := article.New{
				Title: "Article title",
				Body:  "article body",
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			a, err := r.Create(ctx, &na)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if a.ID == 0 {
				t.Error("expected to parse returned id")
			}
		}
	}
}

const (
	articleData = `articles:
  - id: 1
    title: Title
    body: Body
`
)

func TestRepositoryFind(t *testing.T) {
	t.Parallel()
	t.Log("with initialized repository")
	{
		db, teardown := postgresDB(t)
		defer teardown()
		p := polluter.New(polluter.PostgresEngine(db))
		if err := p.Pollute(strings.NewReader(articleData)); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		r := NewRepository(db)

		t.Log("\ttest:0\tshould find record in database")
		{
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			a, err := r.Find(ctx, 1)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if a.Title != "Title" {
				t.Errorf("wrong name scaned: %s", a.Title)
			}
		}

		t.Log("\ttest:1\tshould return not found error")
		{
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if _, err := r.Find(ctx, 333); err != article.ErrNotFound {
				t.Errorf("unexpected error: %v", err)
			}
		}
	}
}
