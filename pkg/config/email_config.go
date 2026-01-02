package config

// Email/SMTP configuration settings
type EmailConfig struct {
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPUser     string `mapstructure:"SMTP_USER"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
	FromEmail    string `mapstructure:"FROM_EMAIL"`
	FromName     string `mapstructure:"FROM_NAME"`
}

// NewEmailConfig creates default email config
func NewEmailConfig() EmailConfig {
	return EmailConfig{
		SMTPPort: 587,
		FromName: "Max Secure",
	}
}

// Validate validates email configuration
func (c *EmailConfig) Validate() error {
	// TODO: Implement validation
	// - Check required fields
	// - Validate email format
	// - Check SMTP port
	return nil
}

// IsConfigured checks if email is properly configured
func (c *EmailConfig) IsConfigured() bool {
	return c.SMTPHost != "" && c.SMTPUser != "" && c.FromEmail != ""
}
