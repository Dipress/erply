package find

import (
	"context"

	"github.com/pkg/errors"
	"github.com/romanyx/erply/internal/article"
)

// Repository is a data access layer.
type Repository interface {
	Find(context.Context, int) (*article.Model, error)
}

// NewService is a factory for service initialization.
func NewService(r Repository) *Service {
	s := Service{
		Repository: r,
	}

	return &s
}

// Service is a use case for use find in a database.
type Service struct {
	Repository
}

// Find finds and returns article.
func (s *Service) Find(ctx context.Context, id int) (*article.Model, error) {
	u, err := s.Repository.Find(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "repository find")
	}

	return u, nil
}
