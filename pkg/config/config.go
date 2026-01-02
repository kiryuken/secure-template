package config

// Configuration loader
// Uses github.com/spf13/viper and github.com/joho/godotenv
// Individual config structs are in separate files:
// - server_config.go
// - database_config.go, redis_config.go
// - aws_config.go
// - auth_config.go
// - email_config.go
// - security_config.go
// - logging_config.go

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	AWS      AWSConfig
	JWT      JWTConfig
	Email    EmailConfig
	Security SecurityConfig
	Logging  LoggingConfig
}

// LoadConfig loads configuration from env and config files
func LoadConfig() (*Config, error) {
	// TODO: Implement with Viper
	return &Config{}, nil
}

// Validate validates configuration
func (c *Config) Validate() error {
	// TODO: Implement validation
	return nil
}
