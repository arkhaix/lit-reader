package common

import "time"

// Cache defines general key-value caching functionality
type Cache interface {
	// Put inserts or updates an entry in the cache.  The entry expires after ttl.
	Put(key string, value string, ttl time.Duration) error

	// Get retrieves a value from the cache.  If the value is not present or is
	// expired, then it returns the empty string and ok=false.  If the value is present,
	// then it returns that value and ok=true.
	Get(key string) (value string, ok bool)

	// Delete removes a value from the cache, if possible.
	Delete(key string)

	// Purge completely clears the cache, if possible.
	Purge()
}
