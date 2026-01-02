package storage

// Storage interface for file storage
// Supports multiple providers (S3, Local, Azure, GCS)

type Storage interface {
	// Upload uploads file to storage
	Upload(filepath, contentType string, data []byte) (string, error)

	// Download downloads file from storage
	Download(filepath string) ([]byte, error)

	// Delete deletes file from storage
	Delete(filepath string) error

	// GetURL gets public/signed URL for file
	GetURL(filepath string, expiry int) (string, error)

	// Exists checks if file exists
	Exists(filepath string) (bool, error)
}

// FileMetadata contains file information
type FileMetadata struct {
	Name        string
	Size        int64
	ContentType string
	URL         string
	StoragePath string
}
