package storage

// Local file storage implementation
// For development/testing or local deployments

type LocalStorage struct {
	// TODO: Add base path
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

// Upload uploads a file to local storage
func (s *LocalStorage) Upload() (string, error) {
	// TODO: Implement local file upload
	return "", nil
}

// Download downloads a file from local storage
func (s *LocalStorage) Download() ([]byte, error) {
	// TODO: Implement local file download
	return nil, nil
}

// Delete deletes a file from local storage
func (s *LocalStorage) Delete() error {
	// TODO: Implement local file delete
	return nil
}

// GetPath gets the full path to a file
func (s *LocalStorage) GetPath() string {
	// TODO: Implement
	return ""
}
