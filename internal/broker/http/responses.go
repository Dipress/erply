package http

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/romanyx/erply/internal/validation"
)

var (
	badRequestBody = messageResponse{
		Message: "bad request",
	}

	internalServerErrorBody = messageResponse{
		Message: "internal server error",
	}

	notFoundBody = messageResponse{
		Message: "not found",
	}
)

type messageResponse struct {
	Message string `json:"message"`
}

func badRequestResponse(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(&badRequestBody); err != nil {
		return errors.Wrap(err, "encoder encode")
	}

	return nil
}

type validationResponse struct {
	Message string            `json:"message"`
	Errors  validation.Errors `json:"errors"`
}

func unprocessabeEntityResponse(w http.ResponseWriter, ers validation.Errors) error {
	w.WriteHeader(http.StatusUnprocessableEntity)

	ver := validationResponse{
		Message: ers.Error(),
		Errors:  ers,
	}

	if err := json.NewEncoder(w).Encode(&ver); err != nil {
		return errors.Wrap(err, "encoder encode")
	}

	return nil
}

func internalServerErrorResponse(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(&internalServerErrorBody); err != nil {
		return errors.Wrap(err, "encoder encode")
	}

	return nil
}

func notFoundResponse(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(&notFoundBody); err != nil {
		return errors.Wrap(err, "encoder encode")
	}

	return nil
}
