package constants

// Rate limiting constants

// Rate limit configurations
const (
	// Global rate limits (requests per second)
	GlobalRateLimitRPS = 100

	// Per-endpoint rate limits
	LoginRateLimitRPS    = 5  // 5 requests per second
	RegisterRateLimitRPS = 2  // 2 requests per second
	UploadRateLimitRPS   = 10 // 10 requests per second

	// Rate limit windows (in seconds)
	RateLimitWindow = 60

	// Max requests per window
	MaxLoginAttempts    = 5
	MaxRegisterAttempts = 3
	MaxUploadAttempts   = 20
)

// Rate limit keys
const (
	RateLimitKeyPrefix = "ratelimit:"
	RateLimitKeyIP     = "ip:"
	RateLimitKeyUser   = "user:"
)
