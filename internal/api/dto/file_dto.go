package dto

// File upload and management DTOs

// FileUploadResponse represents file upload response
type FileUploadResponse struct {
	ID           string `json:"id"`
	Filename     string `json:"filename"`
	OriginalName string `json:"original_name"`
	Size         int64  `json:"size"`
	MimeType     string `json:"mime_type"`
	DownloadURL  string `json:"download_url,omitempty"`
	UploadedAt   string `json:"uploaded_at"`
}

// FileDTO represents file data transfer object
type FileDTO struct {
	ID           string `json:"id"`
	Filename     string `json:"filename"`
	OriginalName string `json:"original_name"`
	Size         int64  `json:"size"`
	MimeType     string `json:"mime_type"`
	StorageType  string `json:"storage_type"`
	IsPublic     bool   `json:"is_public"`
	DownloadURL  string `json:"download_url,omitempty"`
	Checksum     string `json:"checksum"`
	UploadedAt   string `json:"uploaded_at"`
}

// FileListResponse represents file list with pagination
type FileListResponse struct {
	Files      []FileDTO `json:"files"`
	Total      int64     `json:"total"`
	Page       int       `json:"page"`
	PageSize   int       `json:"page_size"`
	TotalPages int       `json:"total_pages"`
}
