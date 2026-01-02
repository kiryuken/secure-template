package errors

// Resource errors

import "net/http"

var (
	ErrNotFound = NewAppError(
		"NOT_FOUND",
		"Resource not found",
		http.StatusNotFound,
	)

	ErrResourceNotFound = NewAppError(
		"RESOURCE_NOT_FOUND",
		"The requested resource was not found",
		http.StatusNotFound,
	)

	ErrResourceAlreadyExists = NewAppError(
		"RESOURCE_ALREADY_EXISTS",
		"Resource already exists",
		http.StatusConflict,
	)

	ErrResourceDeleted = NewAppError(
		"RESOURCE_DELETED",
		"Resource has been deleted",
		http.StatusGone,
	)

	ErrResourceLocked = NewAppError(
		"RESOURCE_LOCKED",
		"Resource is locked",
		http.StatusLocked,
	)
)
