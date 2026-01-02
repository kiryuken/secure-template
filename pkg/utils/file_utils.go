package utils

// File handling utilities
// Helper functions for file operations

import (
	"os"
	"path/filepath"
)

// GetFileHash calculates SHA256 hash of file
func GetFileHash(filepath string) (string, error) {
	// TODO: Implement file hashing
	// - Open file
	// - Calculate SHA256 hash
	// - Return hex string
	return "", nil
}

// GetFileSize gets file size in bytes
func GetFileSize(filepath string) (int64, error) {
	// TODO: Implement
	info, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// GetFileExtension extracts file extension
func GetFileExtension(filename string) string {
	return filepath.Ext(filename)
}

// SanitizeFilename removes unsafe characters from filename
func SanitizeFilename(filename string) string {
	// TODO: Implement filename sanitization
	// - Remove path separators
	// - Remove null bytes
	// - Limit length
	// - Keep extension
	return filename
}

// GenerateUniqueFilename generates unique filename with UUID
func GenerateUniqueFilename(originalFilename string) string {
	// TODO: Implement unique filename generation
	// - Generate UUID
	// - Append original extension
	// Format: {uuid}.{ext}
	return ""
}

// CopyFile copies file from src to dst
func CopyFile(src, dst string) error {
	// TODO: Implement file copy
	// - Open source file
	// - Create destination file
	// - Copy contents
	return nil
}

// EnsureDir creates directory if it doesn't exist
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}
