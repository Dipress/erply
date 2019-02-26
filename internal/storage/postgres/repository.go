package postgres

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/romanyx/erply/internal/article"
)

const (
	driverName = "postgres"
)

// Repository holds crud actions.
type Repository struct {
	db *sql.DB
}

// NewRepository returns ready to work repository.
func NewRepository(db *sql.DB) *Repository {
	r := Repository{
		db: db,
	}

	return &r
}

const createQuery = `INSERT INTO articles (title, body) VALUES ($1, $2) RETURNING id, title, body, created_at, updated_at`

// Create insert article into database.
func (r *Repository) Create(ctx context.Context, f *article.New) (*article.Model, error) {
	var a article.Model
	if err := r.db.QueryRowContext(ctx, createQuery, f.Title, f.Body).
		Scan(&a.ID, &a.Title, &a.Body, &a.CreatedAt, &a.UpdatedAt); err != nil {
		return nil, errors.Wrap(err, "query context scan")
	}

	return &a, nil
}

const findQuery = "SELECT id, title, body, created_at, updated_at FROM articles WHERE id=$1"

// Find finds article by id.
func (r *Repository) Find(ctx context.Context, id int) (*article.Model, error) {
	var a article.Model
	if err := r.db.QueryRowContext(ctx, findQuery, id).
		Scan(&a.ID, &a.Title, &a.Body, &a.CreatedAt, &a.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, article.ErrNotFound
		}
		return nil, errors.Wrap(err, "query row scan")
	}

	return &a, nil
}
