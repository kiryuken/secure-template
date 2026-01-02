package config

// Server configuration settings
type ServerConfig struct {
	Host         string `mapstructure:"SERVER_HOST"`
	Port         int    `mapstructure:"SERVER_PORT"`
	Environment  string `mapstructure:"ENVIRONMENT"`
	ReadTimeout  int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITE_TIMEOUT"`
	IdleTimeout  int    `mapstructure:"IDLE_TIMEOUT"`
}

// NewServerConfig creates default server config
func NewServerConfig() ServerConfig {
	return ServerConfig{
		Host:         "0.0.0.0",
		Port:         8080,
		Environment:  "development",
		ReadTimeout:  15,
		WriteTimeout: 15,
		IdleTimeout:  60,
	}
}

// Validate validates server configuration
func (c *ServerConfig) Validate() error {
	// TODO: Implement validation
	// - Check port range (1-65535)
	// - Validate environment (development, staging, production)
	// - Check timeout values > 0
	return nil
}
