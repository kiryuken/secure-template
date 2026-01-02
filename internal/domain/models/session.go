package models

// Session represents a user session
type Session struct {
	ID           string
	UserID       string
	Token        string
	RefreshToken string
	IPAddress    string
	UserAgent    string
	ExpiresAt    string
	CreatedAt    string
	UpdatedAt    string
}
