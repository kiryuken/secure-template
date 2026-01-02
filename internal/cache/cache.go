package cache

// Cache interface for Redis operations
// Individual use cases in separate files

type Cache interface {
	// Get retrieves value from cache
	Get(key string) (string, error)

	// Set stores value in cache with TTL
	Set(key string, value interface{}, ttl int) error

	// Delete removes value from cache
	Delete(key string) error

	// Exists checks if key exists
	Exists(key string) (bool, error)

	// Increment increments counter
	Increment(key string) (int64, error)

	// Expire sets expiry on existing key
	Expire(key string, ttl int) error
}
