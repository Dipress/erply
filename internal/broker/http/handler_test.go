package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/romanyx/erply/internal/article"
	"github.com/romanyx/erply/internal/create"
	"github.com/romanyx/erply/internal/validation"
)

func TestCreateHandlerCreate(t *testing.T) {
	tt := []struct {
		name        string
		createrFunc func(context.Context, *create.Form) (*article.Model, error)
		code        int
	}{
		{
			name: "ok",
			createrFunc: func(context.Context, *create.Form) (*article.Model, error) {
				return &article.Model{}, nil
			},
			code: http.StatusOK,
		},
		{
			name: "validation error",
			createrFunc: func(context.Context, *create.Form) (*article.Model, error) {
				return nil, make(validation.Errors)
			},
			code: http.StatusUnprocessableEntity,
		},
		{
			name: "internal error",
			createrFunc: func(context.Context, *create.Form) (*article.Model, error) {
				return nil, errors.New("mock error")
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			h := CreateHandler{createrFunc(tc.createrFunc)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://example.com", strings.NewReader("{}"))

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

func TestCreateHandlerFind(t *testing.T) {
	tt := []struct {
		name       string
		finderFunc func(context.Context, int) (*article.Model, error)
		code       int
	}{
		{
			name: "ok",
			finderFunc: func(context.Context, int) (*article.Model, error) {
				return &article.Model{}, nil
			},
			code: http.StatusOK,
		},
		{
			name: "internal error",
			finderFunc: func(context.Context, int) (*article.Model, error) {
				return nil, errors.New("mock error")
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			h := FindHandler{finderFunc(tc.finderFunc)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "http://example.com", nil)
			r = mux.SetURLVars(r, map[string]string{"id": "1"})

			err := h.Handle(w, r)
			if w.Code != tc.code {
				t.Errorf("unexpected code: %d expected %d error: %v", w.Code, tc.code, err)
			}
		})
	}
}

type finderFunc func(context.Context, int) (*article.Model, error)

func (f finderFunc) Find(c context.Context, i int) (*article.Model, error) {
	return f(c, i)
}

type createrFunc func(context.Context, *create.Form) (*article.Model, error)

func (f createrFunc) Create(c context.Context, fm *create.Form) (*article.Model, error) {
	return f(c, fm)
}
