package find

import (
	"context"
	"errors"
	"testing"

	"github.com/romanyx/erply/internal/article"
)

func TestServiceFind(t *testing.T) {
	tt := []struct {
		name           string
		wantErr        bool
		repositoryFunc func(context.Context, int) (*article.Model, error)
	}{
		{
			name: "ok",
			repositoryFunc: func(context.Context, int) (*article.Model, error) {
				return &article.Model{}, nil
			},
		},
		{
			name: "repository error",
			repositoryFunc: func(context.Context, int) (*article.Model, error) {
				return nil, errors.New("mock error")
			},
			wantErr: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := NewService(repositoryFunc(tc.repositoryFunc))
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			_, err := s.Find(ctx, 1)
			if tc.wantErr {
				if err == nil {
					t.Error("expected error")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

type repositoryFunc func(context.Context, int) (*article.Model, error)

func (f repositoryFunc) Find(c context.Context, i int) (*article.Model, error) {
	return f(c, i)
}
