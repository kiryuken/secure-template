package config

// Logging configuration settings
type LoggingConfig struct {
	Level      string `mapstructure:"LOG_LEVEL"`
	OutputPath string `mapstructure:"LOG_OUTPUT_PATH"`
	MaxSize    int    `mapstructure:"LOG_MAX_SIZE"` // MB
	MaxBackups int    `mapstructure:"LOG_MAX_BACKUPS"`
	MaxAge     int    `mapstructure:"LOG_MAX_AGE"` // days
}

// NewLoggingConfig creates default logging config
func NewLoggingConfig() LoggingConfig {
	return LoggingConfig{
		Level:      "info",
		OutputPath: "./logs/app.log",
		MaxSize:    100, // 100MB
		MaxBackups: 5,
		MaxAge:     30, // 30 days
	}
}

// Validate validates logging configuration
func (c *LoggingConfig) Validate() error {
	// TODO: Implement validation
	// - Check valid log level (debug, info, warn, error)
	// - Validate max size > 0
	// - Check max backups >= 0
	// - Validate max age > 0
	return nil
}
