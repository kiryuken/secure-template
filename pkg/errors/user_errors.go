package errors

// User-related errors

import "net/http"

var (
	ErrUserNotFound = NewAppError(
		"USER_NOT_FOUND",
		"User not found",
		http.StatusNotFound,
	)

	ErrUserAlreadyExists = NewAppError(
		"USER_ALREADY_EXISTS",
		"User with this email already exists",
		http.StatusConflict,
	)

	ErrUsernameAlreadyExists = NewAppError(
		"USERNAME_ALREADY_EXISTS",
		"Username already taken",
		http.StatusConflict,
	)

	ErrUserInactive = NewAppError(
		"USER_INACTIVE",
		"User account is inactive",
		http.StatusForbidden,
	)

	ErrUserNotVerified = NewAppError(
		"USER_NOT_VERIFIED",
		"User email is not verified",
		http.StatusForbidden,
	)

	ErrUserSuspended = NewAppError(
		"USER_SUSPENDED",
		"User account has been suspended",
		http.StatusForbidden,
	)

	ErrCannotDeleteSelf = NewAppError(
		"CANNOT_DELETE_SELF",
		"You cannot delete your own account",
		http.StatusBadRequest,
	)
)
