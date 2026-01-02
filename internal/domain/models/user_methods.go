package models

// User domain methods
// Business logic related to User entity

// FullName returns user's full name
func (u *User) FullName() string {
	// TODO: Implement
	return ""
}

// IsAdmin checks if user is admin
func (u *User) IsAdmin() bool {
	// TODO: Implement
	return false
}

// CanAccessResource checks if user can access a resource
func (u *User) CanAccessResource() bool {
	// TODO: Implement
	return false
}

// HasPermission checks if user has specific permission
func (u *User) HasPermission() bool {
	// TODO: Implement
	return false
}

// IsDeleted checks if user is soft-deleted
func (u *User) IsDeleted() bool {
	return u.DeletedAt != nil
}

// Activate activates the user account
func (u *User) Activate() {
	u.IsActive = true
}

// Deactivate deactivates the user account
func (u *User) Deactivate() {
	u.IsActive = false
}

// Verify marks user as verified
func (u *User) Verify() {
	u.IsVerified = true
}
