package config

// Redis configuration settings
type RedisConfig struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     int    `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

// NewRedisConfig creates default Redis config
func NewRedisConfig() RedisConfig {
	return RedisConfig{
		Host: "localhost",
		Port: 6379,
		DB:   0,
	}
}

// Validate validates Redis configuration
func (c *RedisConfig) Validate() error {
	// TODO: Implement validation
	// - Check required fields
	// - Validate port range
	// - Validate DB number (0-15)
	return nil
}

// Address returns Redis connection address
func (c *RedisConfig) Address() string {
	// TODO: Implement
	// Format: host:port
	return ""
}
