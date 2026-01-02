package models

// User base struct
// Only contains core user fields
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

// TableName returns the table name for User
func (User) TableName() string {
	return "users"
}
