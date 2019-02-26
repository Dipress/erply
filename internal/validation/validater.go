package validation

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/romanyx/erply/internal/create"
)

const (
	mismatchMsg   = "mismatch"
	validationMsg = "you have validation errors"
)

// Errors holds validation errors.
type Errors map[string]string

// Error implements error interface.
func (v Errors) Error() string {
	return validationMsg
}

// Create holds create form validations.
type Create struct{}

// Validate valides user form.
func (v *Create) Validate(ctx context.Context, f *create.Form) error {
	ves := make(Errors)

	if err := validation.Validate(f.Title,
		validation.Required,
		validation.Length(1, 50)); err != nil {
		ves["title"] = err.Error()
	}

	if err := validation.Validate(f.Body,
		validation.Required); err != nil {
		ves["body"] = err.Error()
	}

	if len(ves) > 0 {
		return ves
	}

	return nil
}
