package repository

// UserRepository handles user data access
type UserRepository struct {
	// TODO: Add DB connection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create creates a new user
func (r *UserRepository) Create() error {
	// TODO: Implement with Bun ORM
	return nil
}

// GetByID retrieves user by ID
func (r *UserRepository) GetByID() error {
	// TODO: Implement
	return nil
}

// GetByEmail retrieves user by email
func (r *UserRepository) GetByEmail() error {
	// TODO: Implement
	return nil
}

// Update updates user data
func (r *UserRepository) Update() error {
	// TODO: Implement
	return nil
}

// Delete soft-deletes a user
func (r *UserRepository) Delete() error {
	// TODO: Implement
	return nil
}

// List retrieves users with pagination
func (r *UserRepository) List() error {
	// TODO: Implement with cursor/offset pagination
	return nil
}
