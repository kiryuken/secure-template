package constants

// User-related constants

// Default values
const (
	DefaultUserRole   = "user"
	DefaultPageSize   = 20
	MaxPageSize       = 100
	MinPasswordLength = 8
	MaxPasswordLength = 128
	MinUsernameLength = 3
	MaxUsernameLength = 50
	MaxNameLength     = 100
	DefaultAvatarURL  = "/assets/default-avatar.png"
)

// User status
const (
	UserStatusActive    = "active"
	UserStatusInactive  = "inactive"
	UserStatusSuspended = "suspended"
	UserStatusPending   = "pending"
)

// User permissions
const (
	PermissionUserRead   = "user:read"
	PermissionUserWrite  = "user:write"
	PermissionUserDelete = "user:delete"
	PermissionAdminAll   = "admin:*"
)
