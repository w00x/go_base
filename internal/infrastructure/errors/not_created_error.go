package errors

import (
	"go_base/internal/domain/errors"
	"net/http"
)

type NotCreatedError struct {
	errors.BaseError
}

func NewNotCreatedError(message string) *NotCreatedError {
	return &NotCreatedError{BaseError: errors.BaseError{Message: message, Description: "NOT_FOUND", Code: 1002, HttpStatusCode: http.StatusBadRequest}}
}
