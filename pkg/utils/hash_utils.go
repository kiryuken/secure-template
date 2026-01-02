package utils

// Hashing utilities
// Helper functions for hashing and encoding

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Hash calculates SHA256 hash of string
func SHA256Hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// SHA256HashBytes calculates SHA256 hash of bytes
func SHA256HashBytes(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// CompareHash compares hash with data
func CompareHash(hash, data string) bool {
	return SHA256Hash(data) == hash
}

// GenerateRandomString generates random string
func GenerateRandomString(length int) string {
	// TODO: Implement secure random string generation
	// - Use crypto/rand
	// - Base62 encoding (alphanumeric)
	return ""
}

// GenerateSecureToken generates secure token
func GenerateSecureToken(length int) (string, error) {
	// TODO: Implement secure token generation
	// - Use crypto/rand
	// - URL-safe base64 encoding
	return "", nil
}
