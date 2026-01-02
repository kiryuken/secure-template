package auth

// Password hashing and verification
// Uses golang.org/x/crypto/argon2

type PasswordHasher struct {
	// TODO: Add config
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

// HashPassword hashes a password using Argon2
func (h *PasswordHasher) HashPassword() (string, error) {
	// TODO: Implement Argon2 hashing
	return "", nil
}

// VerifyPassword verifies a password against hash
func (h *PasswordHasher) VerifyPassword() (bool, error) {
	// TODO: Implement verification
	return false, nil
}
