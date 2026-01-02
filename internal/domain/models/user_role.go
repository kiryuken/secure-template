package models

// UserRole represents user roles enum
type UserRole string

// Role constants
const (
	RoleAdmin     UserRole = "admin"
	RoleModerator UserRole = "moderator"
	RoleUser      UserRole = "user"
	RoleGuest     UserRole = "guest"
)

// String returns string representation
func (r UserRole) String() string {
	return string(r)
}

// IsValid checks if role is valid
func (r UserRole) IsValid() bool {
	switch r {
	case RoleAdmin, RoleModerator, RoleUser, RoleGuest:
		return true
	}
	return false
}
