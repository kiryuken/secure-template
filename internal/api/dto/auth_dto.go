package dto

// User-related Data Transfer Objects (DTOs)
// Request and Response structures

// LoginRequest represents login request payload
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	MFACode  string `json:"mfa_code,omitempty"`
}

// LoginResponse represents login response
type LoginResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	ExpiresIn    int64   `json:"expires_in"`
	TokenType    string  `json:"token_type"`
	User         UserDTO `json:"user"`
}
