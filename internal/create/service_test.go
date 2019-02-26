package create

import (
	"context"
	"errors"
	"testing"

	"github.com/romanyx/erply/internal/article"
)

func TestServiceCreate(t *testing.T) {
	tt := []struct {
		name           string
		validaterFunc  func(context.Context, *Form) error
		repositoryFunc func(context.Context, *article.New) (*article.Model, error)
		wantErr        bool
	}{
		{
			name: "ok",
			validaterFunc: func(context.Context, *Form) error {
				return nil
			},
			repositoryFunc: func(context.Context, *article.New) (*article.Model, error) {
				return &article.Model{}, nil
			},
		},
		{
			name: "validater error",
			validaterFunc: func(context.Context, *Form) error {
				return errors.New("mock error")
			},
			wantErr: true,
		},
		{
			name: "repository error",
			validaterFunc: func(context.Context, *Form) error {
				return nil
			},
			repositoryFunc: func(context.Context, *article.New) (*article.Model, error) {
				return nil, errors.New("mock error")
			},
			wantErr: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := NewService(repositoryFunc(tc.repositoryFunc), validaterFunc(tc.validaterFunc))
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			var f Form

			_, err := s.Create(ctx, &f)
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

type validaterFunc func(context.Context, *Form) error

func (f validaterFunc) Validate(c context.Context, n *Form) error {
	return f(c, n)
}

type repositoryFunc func(context.Context, *article.New) (*article.Model, error)

func (f repositoryFunc) Create(c context.Context, n *article.New) (*article.Model, error) {
	return f(c, n)
}
