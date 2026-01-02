package dto

// EnableMFARequest represents MFA enable request
type EnableMFARequest struct {
	Password string `json:"password" validate:"required"`
}

// EnableMFAResponse represents MFA enable response
type EnableMFAResponse struct {
	Secret    string `json:"secret"`
	QRCodeURL string `json:"qr_code_url"`
}

// VerifyMFARequest represents MFA verification request
type VerifyMFARequest struct {
	Code string `json:"code" validate:"required,len=6"`
}

// VerifyMFAResponse represents MFA verification response
type VerifyMFAResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
