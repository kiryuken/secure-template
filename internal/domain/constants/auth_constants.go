package constants

// Authentication-related constants

// Token expiration times (in seconds)
const (
	AccessTokenExpiry  = 15 * 60          // 15 minutes
	RefreshTokenExpiry = 7 * 24 * 60 * 60 // 7 days
	ResetTokenExpiry   = 1 * 60 * 60      // 1 hour
	VerifyTokenExpiry  = 24 * 60 * 60     // 24 hours
)

// JWT claims
const (
	ClaimUserID      = "user_id"
	ClaimEmail       = "email"
	ClaimRole        = "role"
	ClaimPermissions = "permissions"
	ClaimSessionID   = "session_id"
)

// Session constants
const (
	SessionKeyPrefix     = "session:"
	SessionMaxAge        = 7 * 24 * 60 * 60 // 7 days
	MaxConcurrentSession = 5
)

// MFA constants
const (
	MFAIssuer     = "MaxSecure"
	MFACodeLength = 6
	MFAWindowSize = 1
)
