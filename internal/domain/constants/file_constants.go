package constants

// File-related constants

// File size limits (in bytes)
const (
	MaxFileSize     = 10 * 1024 * 1024  // 10 MB
	MaxImageSize    = 5 * 1024 * 1024   // 5 MB
	MaxDocumentSize = 20 * 1024 * 1024  // 20 MB
	MaxVideoSize    = 100 * 1024 * 1024 // 100 MB
)

// Allowed file types
var (
	AllowedImageTypes = []string{
		"image/jpeg",
		"image/png",
		"image/gif",
		"image/webp",
	}

	AllowedDocumentTypes = []string{
		"application/pdf",
		"application/msword",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"text/plain",
	}

	AllowedVideoTypes = []string{
		"video/mp4",
		"video/mpeg",
		"video/webm",
	}
)

// File storage paths
const (
	StoragePathAvatars   = "avatars"
	StoragePathDocuments = "documents"
	StoragePathUploads   = "uploads"
	StoragePathTemp      = "temp"
)
