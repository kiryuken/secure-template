package constants

// Cache key constants

// Cache key prefixes
const (
	CacheKeyUser       = "user:"
	CacheKeySession    = "session:"
	CacheKeyToken      = "token:"
	CacheKeyRateLimit  = "ratelimit:"
	CacheKeyBlacklist  = "blacklist:"
	CacheKeyPermission = "permission:"
)

// Cache TTL (in seconds)
const (
	CacheTTLShort  = 5 * 60       // 5 minutes
	CacheTTLMedium = 30 * 60      // 30 minutes
	CacheTTLLong   = 60 * 60      // 1 hour
	CacheTTLDay    = 24 * 60 * 60 // 24 hours
)
