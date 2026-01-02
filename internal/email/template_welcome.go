package email

// Welcome email template
// Sent when user registers successfully

type WelcomeEmail struct {
	service EmailService
}

// NewWelcomeEmail creates welcome email sender
func NewWelcomeEmail(service EmailService) *WelcomeEmail {
	return &WelcomeEmail{service: service}
}

// Send sends welcome email to new user
func (e *WelcomeEmail) Send( /* toEmail, userName string */ ) error {
	// TODO: Implement welcome email
	// Subject: "Welcome to Max Secure!"
	// Content:
	// - Welcome message
	// - Getting started guide
	// - Important links (docs, support)
	// - Contact information

	_ = "Welcome to Max Secure!"                    // subject
	_ = `<!-- TODO: Design HTML email template -->` // htmlBody

	return nil
}
