package logger

// Log rotation configuration
// Uses gopkg.in/natefinch/lumberjack.v2

type LogRotator struct {
	// TODO: Add lumberjack config
}

func NewLogRotator() *LogRotator {
	return &LogRotator{}
}

// Setup configures log rotation
func (r *LogRotator) Setup() error {
	// TODO: Implement
	// - Max size
	// - Max backups
	// - Max age
	// - Compress
	return nil
}
