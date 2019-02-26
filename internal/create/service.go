package create

import (
	"context"

	"github.com/pkg/errors"
	"github.com/romanyx/erply/internal/article"
)

// Repository is a data access layer.
type Repository interface {
	Create(context.Context, *article.New) (*article.Model, error)
}

// Validater allows to validate article.
type Validater interface {
	Validate(context.Context, *Form) error
}

// NewService is a factory for service initialization.
func NewService(r Repository, v Validater) *Service {
	s := Service{
		Repository: r,
		Validater:  v,
	}

	return &s
}

// Service is a use case for article validation and creation.
type Service struct {
	Repository
	Validater
}

// Form is a create request.
type Form struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Create validates and creates article.
func (s *Service) Create(ctx context.Context, f *Form) (*article.Model, error) {
	if err := s.Validate(ctx, f); err != nil {
		return nil, errors.Wrap(err, "validater validate")
	}

	na := article.New{
		Title: f.Title,
		Body:  f.Body,
	}

	a, err := s.Repository.Create(ctx, &na)
	if err != nil {
		return nil, errors.Wrap(err, "repository create")
	}

	return a, nil
}
