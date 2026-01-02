package cache

// Token blacklist caching
// Stores revoked JWT tokens in Redis

type TokenCache struct {
	// TODO: Add dependencies
	// redis Cache
}

// NewTokenCache creates token cache
func NewTokenCache( /* redis Cache */ ) *TokenCache {
	return &TokenCache{
		// TODO: Initialize
	}
}

// BlacklistToken adds token to blacklist
func (c *TokenCache) BlacklistToken( /* token string, expiry int */ ) error {
	// TODO: Implement
	// - Hash token (SHA256)
	// - Store with key: blacklist:{hash}
	// - Set TTL to match token expiry
	return nil
}

// IsBlacklisted checks if token is blacklisted
func (c *TokenCache) IsBlacklisted( /* token string */ ) (bool, error) {
	// TODO: Implement
	// - Hash token
	// - Check if key exists: blacklist:{hash}
	return false, nil
}

// CleanupExpired removes expired blacklisted tokens
func (c *TokenCache) CleanupExpired() error {
	// TODO: Implement
	// Note: Redis TTL handles this automatically
	// This method for manual cleanup if needed
	return nil
}
