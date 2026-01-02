package models

// User represents a user in the system
type User struct {
	ID           string
	Email        string
	Username     string
	PasswordHash string
	FirstName    string
	LastName     string
	IsActive     bool
	IsVerified   bool
	MFAEnabled   bool
	MFASecret    string
	Role         string
	Permissions  []string
	AvatarURL    string
	LastLoginAt  string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    *string
}

// UserRole represents user roles
type UserRole string

const (
	RoleAdmin     UserRole = "admin"
	RoleModerator UserRole = "moderator"
	RoleUser      UserRole = "user"
)
