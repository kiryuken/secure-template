package config

// Database configuration settings
type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_NAME"`
	SSLMode  string `mapstructure:"DB_SSL_MODE"`
	MaxConns int    `mapstructure:"DB_MAX_CONNS"`
	MinConns int    `mapstructure:"DB_MIN_CONNS"`
}

// NewDatabaseConfig creates default database config
func NewDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		SSLMode:  "disable",
		MaxConns: 25,
		MinConns: 5,
	}
}

// Validate validates database configuration
func (c *DatabaseConfig) Validate() error {
	// TODO: Implement validation
	// - Check required fields (Host, User, Database)
	// - Validate port range
	// - Check connection pool settings
	return nil
}

// DSN returns PostgreSQL connection string
func (c *DatabaseConfig) DSN() string {
	// TODO: Implement DSN builder
	// Format: postgres://user:password@host:port/database?sslmode=disable
	return ""
}
