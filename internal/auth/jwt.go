package auth

// JWT token generation and validation
// Uses github.com/golang-jwt/jwt/v5

type JWTManager struct {
	// TODO: Add config
}

func NewJWTManager() *JWTManager {
	return &JWTManager{}
}

// GenerateToken generates a JWT token
func (m *JWTManager) GenerateToken() (string, error) {
	// TODO: Implement JWT generation
	return "", nil
}

// GenerateRefreshToken generates a refresh token
func (m *JWTManager) GenerateRefreshToken() (string, error) {
	// TODO: Implement
	return "", nil
}

// ValidateToken validates and parses JWT token
func (m *JWTManager) ValidateToken() error {
	// TODO: Implement JWT validation
	return nil
}

// ExtractClaims extracts claims from JWT token
func (m *JWTManager) ExtractClaims() error {
	// TODO: Implement
	return nil
}

// RevokeToken revokes a token
func (m *JWTManager) RevokeToken() error {
	// TODO: Implement (use Redis blacklist)
	return nil
}
