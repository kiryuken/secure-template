package errors

// Custom error types and error handling
// Uses github.com/pkg/errors

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	// TODO: Implement
	return ""
}

// NewAppError creates a new application error
func NewAppError() *AppError {
	return &AppError{}
}

// Wrap wraps an error with context
func Wrap() error {
	// TODO: Implement with pkg/errors
	return nil
}

// Common error types
var (
	ErrNotFound           = &AppError{Code: "NOT_FOUND", Message: "Resource not found"}
	ErrUnauthorized       = &AppError{Code: "UNAUTHORIZED", Message: "Unauthorized"}
	ErrForbidden          = &AppError{Code: "FORBIDDEN", Message: "Forbidden"}
	ErrBadRequest         = &AppError{Code: "BAD_REQUEST", Message: "Bad request"}
	ErrInternalServer     = &AppError{Code: "INTERNAL_SERVER", Message: "Internal server error"}
	ErrInvalidCredentials = &AppError{Code: "INVALID_CREDENTIALS", Message: "Invalid credentials"}
	ErrInvalidToken       = &AppError{Code: "INVALID_TOKEN", Message: "Invalid token"}
	ErrExpiredToken       = &AppError{Code: "EXPIRED_TOKEN", Message: "Token expired"}
	ErrInvalidMFA         = &AppError{Code: "INVALID_MFA", Message: "Invalid MFA code"}
)
