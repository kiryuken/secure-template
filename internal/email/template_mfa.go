package email

// MFA notification email template
// Sent when MFA is enabled/disabled

type MFAEmail struct {
	service EmailService
}

// NewMFAEmail creates MFA email sender
func NewMFAEmail(service EmailService) *MFAEmail {
	return &MFAEmail{service: service}
}

// SendMFAEnabled sends notification when MFA is enabled
func (e *MFAEmail) SendMFAEnabled( /* toEmail, timestamp string */ ) error {
	// TODO: Implement MFA enabled notification
	// Subject: "Two-Factor Authentication Enabled"
	// Content:
	// - Confirmation of MFA setup
	// - Timestamp of change
	// - Backup codes reminder
	// - What to do if not authorized

	_ = "Two-Factor Authentication Enabled"         // subject
	_ = `<!-- TODO: Design HTML email template -->` // htmlBody

	return nil
}

// SendMFADisabled sends notification when MFA is disabled
func (e *MFAEmail) SendMFADisabled( /* toEmail, timestamp string */ ) error {
	// TODO: Implement MFA disabled notification
	// Subject: "Two-Factor Authentication Disabled"
	// Content:
	// - Warning that MFA was disabled
	// - Timestamp of change
	// - Security impact warning
	// - What to do if not authorized

	_ = "Two-Factor Authentication Disabled"        // subject
	_ = `<!-- TODO: Design HTML email template -->` // htmlBody

	return nil
}
