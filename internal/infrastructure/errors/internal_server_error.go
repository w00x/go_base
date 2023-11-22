package errors

import (
	"go_base/internal/domain/errors"
	"net/http"
)

type InternalServerError struct {
	errors.BaseError
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{BaseError: errors.BaseError{Message: message, Description: "INTERNAL_SERVER_ERROR", Code: 1001, HttpStatusCode: http.StatusInternalServerError}}
}
