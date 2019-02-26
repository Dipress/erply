package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/romanyx/erply/internal/article"
	"github.com/romanyx/erply/internal/create"
	"github.com/romanyx/erply/internal/find"
	"github.com/romanyx/erply/internal/validation"
)

const (
	timeout = 30 * time.Second
)

// Repository is a data access layer.
type Repository interface {
	Create(context.Context, *article.New) (*article.Model, error)
	Find(context.Context, int) (*article.Model, error)
}

// NewServer prepares http server.
func NewServer(addr, token string, repo Repository) *http.Server {
	mux := mux.NewRouter()

	createService := create.NewService(repo, &validation.Create{})
	createHandler := CreateHandler{
		Creater: createService,
	}

	findService := find.NewService(repo)
	findHandler := FindHandler{
		Finder: findService,
	}

	mux.HandleFunc("/articles", Secure(httpHandler{
		Handler: &createHandler,
	}, token).ServeHTTP).Methods("POST")

	mux.HandleFunc("/articles/{id}", Secure(httpHandler{
		Handler: &findHandler,
	}, token).ServeHTTP).Methods("GET")

	s := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	return &s
}
