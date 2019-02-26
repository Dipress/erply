package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/romanyx/erply/internal/article"
	"github.com/romanyx/erply/internal/create"
	"github.com/romanyx/erply/internal/validation"
)

// Creater abstraction for create service.
type Creater interface {
	Create(context.Context, *create.Form) (*article.Model, error)
}

// Finder abstraction for find service.
type Finder interface {
	Find(context.Context, int) (*article.Model, error)
}

// Handler allows to handle requests.
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}

// CreateHandler for create requests.
type CreateHandler struct {
	Creater
}

// Handle handles registration requests and return error.
func (h *CreateHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	var f create.Form
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return errors.Wrap(err, "decoder decode")
	}

	u, err := h.Create(r.Context(), &f)
	if err != nil {
		switch v := errors.Cause(err).(type) {
		case validation.Errors:
			w.WriteHeader(http.StatusUnprocessableEntity)
			if err := json.NewEncoder(w).Encode(v); err != nil {
				return errors.Wrap(err, "validation encoder encode")
			}
			return nil
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return errors.Wrap(err, "registrater registrate")
		}
	}

	if err := json.NewEncoder(w).Encode(&u); err != nil {
		return errors.Wrap(err, "encoder encode")
	}

	return nil
}

// FindHandler for create requests.
type FindHandler struct {
	Finder
}

// Handle handles registration requests and return error.
func (h *FindHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return errors.Wrapf(err, "convert id query param to int: %v", err)
	}

	u, err := h.Find(r.Context(), id)
	if err != nil {
		switch errors.Cause(err) {
		case article.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			return nil
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return errors.Wrap(err, "finder find")
		}
	}

	if err := json.NewEncoder(w).Encode(&u); err != nil {
		return errors.Wrap(err, "encoder encode")
	}

	return nil
}

// httpHandler allows to implement ServeHTTP for Handler.
type httpHandler struct {
	Handler
}

// ServeHTTP implements http.Handler.
func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.Handle(w, r); err != nil {
		log.Printf("serve http: %v\n", err)
	}
}
