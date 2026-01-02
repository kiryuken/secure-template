package models

// PasswordResetToken represents a password reset token
type PasswordResetToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt string
	UsedAt    *string
	CreatedAt string
}

// EmailVerificationToken represents an email verification token
type EmailVerificationToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt string
	UsedAt    *string
	CreatedAt string
}
