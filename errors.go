package apperror

import (
	"fmt"
	"net/http"
)

// AppError represents a custom application error
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new AppError
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// HTTPStatus returns the HTTP status code for the error
func (e *AppError) HTTPStatus() int {
	return e.Code
}

// Predefined errors
var (
	ErrNotFound       = NewAppError(http.StatusNotFound, "Resource not found", nil)
	ErrBadRequest     = NewAppError(http.StatusBadRequest, "Bad request", nil)
	ErrUnauthorized   = NewAppError(http.StatusUnauthorized, "Unauthorized", nil)
	ErrForbidden      = NewAppError(http.StatusForbidden, "Forbidden", nil)
	ErrConflict       = NewAppError(http.StatusConflict, "Conflict", nil)
	ErrInternalServer = NewAppError(http.StatusInternalServerError, "Internal server error", nil)
)

// User-specific errors
var (
	ErrUserNotFound    = NewAppError(http.StatusNotFound, "User not found", nil)
	ErrUserExists      = NewAppError(http.StatusConflict, "User already exists", nil)
	ErrInvalidEmail    = NewAppError(http.StatusBadRequest, "Invalid email format", nil)
	ErrInvalidPassword = NewAppError(http.StatusBadRequest, "Invalid password", nil)
	ErrInvalidUsername = NewAppError(http.StatusBadRequest, "Invalid username", nil)
	ErrUserNotVerified = NewAppError(http.StatusForbidden, "User not verified", nil)
)
