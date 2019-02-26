package article

import (
	"errors"
	"time"
)

var (
	// ErrNotFound raises when article not found in database.
	ErrNotFound = errors.New("article not found")
)

// Model is a article representation.
type Model struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// New represents data required to create article.
type New struct {
	Title string
	Body  string
}
