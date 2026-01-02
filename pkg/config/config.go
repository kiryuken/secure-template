package config

// Configuration loader
// Uses github.com/spf13/viper and github.com/joho/godotenv

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

type ServerConfig struct {
	Host         string
	Port         int
	Environment  string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
	MaxConns int
	MinConns int
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type AWSConfig struct {
	Region          string
	S3Bucket        string
	AccessKeyID     string
	SecretAccessKey string
}

type JWTConfig struct {
	Secret          string
	AccessTokenTTL  int
	RefreshTokenTTL int
	Issuer          string
}

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
	FromEmail    string
	FromName     string
}

type SecurityConfig struct {
	AllowedOrigins []string
	RateLimitRPS   int
	MaxUploadSize  int64
	PasswordMinLen int
}

type LoggingConfig struct {
	Level      string
	OutputPath string
	MaxSize    int
	MaxBackups int
	MaxAge     int
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
