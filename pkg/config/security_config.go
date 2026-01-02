package config

// Security configuration settings
type SecurityConfig struct {
	AllowedOrigins []string `mapstructure:"ALLOWED_ORIGINS"`
	RateLimitRPS   int      `mapstructure:"RATE_LIMIT_RPS"`
	MaxUploadSize  int64    `mapstructure:"MAX_UPLOAD_SIZE"` // bytes
	PasswordMinLen int      `mapstructure:"PASSWORD_MIN_LEN"`
	EncryptionKey  string   `mapstructure:"ENCRYPTION_KEY"`
}

// NewSecurityConfig creates default security config
func NewSecurityConfig() SecurityConfig {
	return SecurityConfig{
		AllowedOrigins: []string{"http://localhost:3000"},
		RateLimitRPS:   100,
		MaxUploadSize:  10 * 1024 * 1024, // 10MB
		PasswordMinLen: 8,
	}
}

// Validate validates security configuration
func (c *SecurityConfig) Validate() error {
	// TODO: Implement validation
	// - Check encryption key length (32 bytes)
	// - Validate rate limit > 0
	// - Check max upload size > 0
	// - Validate password min length >= 8
	return nil
}
