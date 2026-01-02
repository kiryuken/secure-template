package logger

// Logger wrapper using zap
// Uses go.uber.org/zap

type Logger struct {
	// TODO: Add zap logger
}

func NewLogger() *Logger {
	return &Logger{}
}

// Debug logs debug message
func (l *Logger) Debug() {
	// TODO: Implement
}

// Info logs info message
func (l *Logger) Info() {
	// TODO: Implement
}

// Warn logs warning message
func (l *Logger) Warn() {
	// TODO: Implement
}

// Error logs error message
func (l *Logger) Error() {
	// TODO: Implement
}

// Fatal logs fatal message and exits
func (l *Logger) Fatal() {
	// TODO: Implement
}

// WithFields adds fields to logger context
func (l *Logger) WithFields() *Logger {
	// TODO: Implement
	return l
}

// WithError adds error to logger context
func (l *Logger) WithError() *Logger {
	// TODO: Implement
	return l
}
