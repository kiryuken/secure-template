package email

// Password reset email template
// Sent when user requests password reset

type PasswordResetEmail struct {
	service EmailService
}

// NewPasswordResetEmail creates password reset email sender
func NewPasswordResetEmail(service EmailService) *PasswordResetEmail {
	return &PasswordResetEmail{service: service}
}

// Send sends password reset link
func (e *PasswordResetEmail) Send( /* toEmail, token, expiryMinutes string */ ) error {
	// TODO: Implement password reset email
	// Subject: "Reset Your Password"
	// Content:
	// - Password reset link with token
	// - Link expiry time (1 hour)
	// - Security warning
	// - What to do if not requested
	// - Contact support if suspicious

	_ = "Reset Your Password"                       // subject
	_ = `<!-- TODO: Design HTML email template -->` // htmlBody

	return nil
}
