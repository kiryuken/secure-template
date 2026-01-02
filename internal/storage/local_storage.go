package storage

// Local filesystem storage implementation
// For development/testing or small deployments

import "os"

type LocalStorage struct {
	// TODO: Add dependencies
	// basePath string
	// baseURL  string
	// logger   logger.Logger
}

// NewLocalStorage creates local storage instance
func NewLocalStorage( /* basePath, baseURL, logger */ ) (*LocalStorage, error) {
	// TODO: Implement
	// - Validate base path exists or create
	// - Set up directory structure
	return &LocalStorage{}, nil
}

// Upload uploads file to local filesystem
func (s *LocalStorage) Upload(filepath, contentType string, data []byte) (string, error) {
	// TODO: Implement local file upload
	// - Generate unique filename
	// - Create directory structure if needed
	// - Write file to disk
	// - Return file path
	return "", nil
}

// Download downloads file from local filesystem
func (s *LocalStorage) Download(filepath string) ([]byte, error) {
	// TODO: Implement local file read
	// - Read file from disk
	// - Return bytes
	return os.ReadFile(filepath)
}

// Delete deletes file from local filesystem
func (s *LocalStorage) Delete(filepath string) error {
	// TODO: Implement file deletion
	// - Delete file from disk
	return os.Remove(filepath)
}

// GetURL gets URL for locally stored file
func (s *LocalStorage) GetURL(filepath string, expiry int) (string, error) {
	// TODO: Implement URL generation
	// - Return baseURL + filepath
	// Note: expiry not applicable for local storage
	return "", nil
}

// Exists checks if file exists in local filesystem
func (s *LocalStorage) Exists(filepath string) (bool, error) {
	// TODO: Implement existence check
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}
