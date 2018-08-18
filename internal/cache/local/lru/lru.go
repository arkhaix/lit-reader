package lru

import (
	"errors"
	"time"

	hashilru "github.com/hashicorp/golang-lru"

	common "github.com/arkhaix/lit-reader/common/cache"
)

// cache implements cache.Cache
type cache struct {
	lru *hashilru.Cache
}

type entry struct {
	value  string
	expiry time.Time
}

// NewCache returns a new LRU cache which will automatically evict the oldest
// Put inserts or updates an entry in the cache.  The entry expires after ttl.
// entry upon inserting new entries beyond maxEntries.
func NewCache(maxEntries int) (common.Cache, error) {
	hashiCache, err := hashilru.New(maxEntries)
	if err != nil {
		return nil, err
	}
	res := cache{
		lru: hashiCache,
	}
	return common.Cache(&res), nil
}

// Put inserts or updates an entry in the cache.  The entry expires after ttl.
func (c *cache) Put(key string, value string, ttl time.Duration) error {
	e := entry{
		value:  value,
		expiry: time.Now().UTC().Add(ttl),
	}

	if e.expiry.Before(time.Now().UTC()) {
		return errors.New("ttl must not be negative")
	}

	c.lru.Add(key, e)
	return nil
}

// Get retrieves a value from the cache.  If the value is not present or is
// expired, then it returns the empty string and ok=false.  If the value is present,
// then it returns that value and ok=true.
func (c *cache) Get(key string) (string, bool) {
	// Retrieve interface from the cache
	entryInterface, ok := c.lru.Get(key)
	if !ok {
		return "", false
	}

	// Convert interface to concrete
	e, ok := entryInterface.(entry)
	if !ok {
		c.lru.Remove(key)
		return "", false
	}

	// Check expiry
	if e.expiry.Before(time.Now().UTC()) {
		c.lru.Remove(key)
		return "", false
	}

	return e.value, true
}

// Delete removes a value from the cache, if possible.
func (c *cache) Delete(key string) {
	c.lru.Remove(key)
}

// Purge completely clears the cache, if possible.
func (c *cache) Purge() {
	c.lru.Purge()
}
