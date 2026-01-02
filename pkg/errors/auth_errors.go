package errors

// Authentication and authorization errors

import "net/http"

var (
	// Authentication errors
	ErrInvalidCredentials = NewAppError(
		"INVALID_CREDENTIALS",
		"Invalid email or password",
		http.StatusUnauthorized,
	)

	ErrInvalidToken = NewAppError(
		"INVALID_TOKEN",
		"Invalid or malformed token",
		http.StatusUnauthorized,
	)

	ErrExpiredToken = NewAppError(
		"EXPIRED_TOKEN",
		"Token has expired",
		http.StatusUnauthorized,
	)

	ErrInvalidMFA = NewAppError(
		"INVALID_MFA",
		"Invalid MFA code",
		http.StatusUnauthorized,
	)

	ErrMFARequired = NewAppError(
		"MFA_REQUIRED",
		"MFA verification required",
		http.StatusUnauthorized,
	)

	// Authorization errors
	ErrUnauthorized = NewAppError(
		"UNAUTHORIZED",
		"You are not authorized to perform this action",
		http.StatusUnauthorized,
	)

	ErrForbidden = NewAppError(
		"FORBIDDEN",
		"Access to this resource is forbidden",
		http.StatusForbidden,
	)

	ErrInsufficientPermissions = NewAppError(
		"INSUFFICIENT_PERMISSIONS",
		"You don't have sufficient permissions",
		http.StatusForbidden,
	)
)
