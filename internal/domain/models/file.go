package models

// File represents an uploaded file
type File struct {
	ID           string
	UserID       string
	Filename     string
	OriginalName string
	MimeType     string
	Size         int64
	StoragePath  string
	StorageType  string // "local", "s3"
	IsPublic     bool
	DownloadURL  string
	Checksum     string
	UploadedAt   string
	DeletedAt    *string
}

// StorageType represents storage backend types
type StorageType string

const (
	StorageLocal StorageType = "local"
	StorageS3    StorageType = "s3"
)
