package validator

// HTML/input sanitization using bluemonday
// Prevents XSS attacks

type Sanitizer struct {
	// TODO: Add bluemonday policy
}

func NewSanitizer() *Sanitizer {
	return &Sanitizer{}
}

// SanitizeHTML sanitizes HTML content
func (s *Sanitizer) SanitizeHTML() string {
	// TODO: Implement with bluemonday
	return ""
}

// SanitizeUserInput sanitizes general user input
func (s *Sanitizer) SanitizeUserInput() string {
	// TODO: Implement
	return ""
}

// StripTags removes all HTML tags
func (s *Sanitizer) StripTags() string {
	// TODO: Implement
	return ""
}
