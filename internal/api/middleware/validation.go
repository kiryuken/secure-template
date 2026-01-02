package middleware

// Request validation middleware
// Validates and sanitizes input data

type ValidationMiddleware struct {
	// TODO: Add validator
}

func NewValidationMiddleware() *ValidationMiddleware {
	return &ValidationMiddleware{}
}

// ValidateRequest validates request body
func (m *ValidationMiddleware) ValidateRequest() {
	// TODO: Implement with go-playground/validator
}

// SanitizeInput sanitizes HTML/user input
func (m *ValidationMiddleware) SanitizeInput() {
	// TODO: Implement with microcosm-cc/bluemonday
}

// ValidateContentType checks request content-type
func (m *ValidationMiddleware) ValidateContentType() {
	// TODO: Implement
}
