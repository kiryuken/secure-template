package config

// AWS configuration settings
type AWSConfig struct {
	Region          string `mapstructure:"AWS_REGION"`
	S3Bucket        string `mapstructure:"S3_BUCKET"`
	AccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

// NewAWSConfig creates default AWS config
func NewAWSConfig() AWSConfig {
	return AWSConfig{
		Region: "us-east-1",
	}
}

// Validate validates AWS configuration
func (c *AWSConfig) Validate() error {
	// TODO: Implement validation
	// - Check required fields for S3 (Region, Bucket, Credentials)
	// - Validate region format
	return nil
}

// IsConfigured checks if AWS is properly configured
func (c *AWSConfig) IsConfigured() bool {
	return c.Region != "" && c.S3Bucket != "" &&
		c.AccessKeyID != "" && c.SecretAccessKey != ""
}
