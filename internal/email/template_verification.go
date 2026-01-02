package email

// Email verification template
// Sent to verify user's email address

type VerificationEmail struct {
	service EmailService
}

// NewVerificationEmail creates verification email sender
func NewVerificationEmail(service EmailService) *VerificationEmail {
	return &VerificationEmail{service: service}
}

// Send sends email verification link
func (e *VerificationEmail) Send( /* toEmail, token, expiryMinutes string */ ) error {
	// TODO: Implement verification email
	// Subject: "Verify Your Email Address"
	// Content:
	// - Verification link with token
	// - Link expiry time
	// - Security notice
	// - What to do if not requested

	_ = "Verify Your Email Address"                 // subject
	_ = `<!-- TODO: Design HTML email template -->` // htmlBody

	return nil
}
