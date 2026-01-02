package config

// JWT authentication configuration
type JWTConfig struct {
	Secret          string `mapstructure:"JWT_SECRET"`
	AccessTokenTTL  int    `mapstructure:"JWT_ACCESS_TTL"`  // minutes
	RefreshTokenTTL int    `mapstructure:"JWT_REFRESH_TTL"` // days
	Issuer          string `mapstructure:"JWT_ISSUER"`
}

// NewJWTConfig creates default JWT config
func NewJWTConfig() JWTConfig {
	return JWTConfig{
		AccessTokenTTL:  15, // 15 minutes
		RefreshTokenTTL: 7,  // 7 days
		Issuer:          "max-secure",
	}
}

// Validate validates JWT configuration
func (c *JWTConfig) Validate() error {
	// TODO: Implement validation
	// - Check JWT secret length (min 32 characters)
	// - Validate TTL values > 0
	// - Check issuer is set
	return nil
}
