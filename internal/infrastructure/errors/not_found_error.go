package errors

import (
	"go_base/internal/domain/errors"
	"net/http"
)

type NotFoundError struct {
	errors.BaseError
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{BaseError: errors.BaseError{Message: message, Description: "NOT_FOUND", Code: 1000, HttpStatusCode: http.StatusNotFound}}
}
