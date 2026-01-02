package errors

// Validation errors

import "net/http"

var (
	ErrBadRequest = NewAppError(
		"BAD_REQUEST",
		"Invalid request data",
		http.StatusBadRequest,
	)

	ErrInvalidInput = NewAppError(
		"INVALID_INPUT",
		"Invalid input data",
		http.StatusBadRequest,
	)

	ErrValidationFailed = NewAppError(
		"VALIDATION_FAILED",
		"Validation failed",
		http.StatusBadRequest,
	)

	ErrInvalidEmail = NewAppError(
		"INVALID_EMAIL",
		"Invalid email format",
		http.StatusBadRequest,
	)

	ErrWeakPassword = NewAppError(
		"WEAK_PASSWORD",
		"Password is too weak",
		http.StatusBadRequest,
	)

	ErrPasswordTooShort = NewAppError(
		"PASSWORD_TOO_SHORT",
		"Password must be at least 8 characters",
		http.StatusBadRequest,
	)

	ErrInvalidFileType = NewAppError(
		"INVALID_FILE_TYPE",
		"File type not allowed",
		http.StatusBadRequest,
	)

	ErrFileTooLarge = NewAppError(
		"FILE_TOO_LARGE",
		"File size exceeds maximum limit",
		http.StatusBadRequest,
	)
)
