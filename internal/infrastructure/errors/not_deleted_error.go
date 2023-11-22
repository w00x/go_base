package errors

import (
	"go_base/internal/domain/errors"
	"net/http"
)

type NotDeletedError struct {
	errors.BaseError
}

func NewNotDeletedError(message string) *NotDeletedError {
	return &NotDeletedError{BaseError: errors.BaseError{Message: message, Description: "NOT_FOUND", Code: 1003, HttpStatusCode: http.StatusBadRequest}}
}
