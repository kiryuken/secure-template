package errors

// System and server errors

import "net/http"

var (
	ErrInternalServer = NewAppError(
		"INTERNAL_SERVER_ERROR",
		"An internal server error occurred",
		http.StatusInternalServerError,
	)

	ErrDatabaseError = NewAppError(
		"DATABASE_ERROR",
		"Database operation failed",
		http.StatusInternalServerError,
	)

	ErrCacheError = NewAppError(
		"CACHE_ERROR",
		"Cache operation failed",
		http.StatusInternalServerError,
	)

	ErrStorageError = NewAppError(
		"STORAGE_ERROR",
		"Storage operation failed",
		http.StatusInternalServerError,
	)

	ErrServiceUnavailable = NewAppError(
		"SERVICE_UNAVAILABLE",
		"Service temporarily unavailable",
		http.StatusServiceUnavailable,
	)

	ErrTimeout = NewAppError(
		"TIMEOUT",
		"Request timeout",
		http.StatusRequestTimeout,
	)

	ErrRateLimitExceeded = NewAppError(
		"RATE_LIMIT_EXCEEDED",
		"Rate limit exceeded",
		http.StatusTooManyRequests,
	)
)
