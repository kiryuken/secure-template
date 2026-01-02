package dto

// User DTO for API responses
type UserDTO struct {
	ID          string   `json:"id"`
	Email       string   `json:"email"`
	Username    string   `json:"username"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	FullName    string   `json:"full_name"`
	IsActive    bool     `json:"is_active"`
	IsVerified  bool     `json:"is_verified"`
	MFAEnabled  bool     `json:"mfa_enabled"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	AvatarURL   string   `json:"avatar_url,omitempty"`
	LastLoginAt string   `json:"last_login_at,omitempty"`
	CreatedAt   string   `json:"created_at"`
}

// CreateUserRequest represents create user request
type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,min=3,max=50"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required,max=100"`
	LastName  string `json:"last_name" validate:"required,max=100"`
	Role      string `json:"role" validate:"omitempty,oneof=admin moderator user"`
}

// UpdateUserRequest represents update user request
type UpdateUserRequest struct {
	FirstName string `json:"first_name,omitempty" validate:"omitempty,max=100"`
	LastName  string `json:"last_name,omitempty" validate:"omitempty,max=100"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

// UserListResponse represents user list response with pagination
type UserListResponse struct {
	Users      []UserDTO `json:"users"`
	Total      int64     `json:"total"`
	Page       int       `json:"page"`
	PageSize   int       `json:"page_size"`
	TotalPages int       `json:"total_pages"`
}
