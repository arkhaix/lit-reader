package lru_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/arkhaix/lit-reader/common/cache"
	. "github.com/arkhaix/lit-reader/common/cache/local/lru"
)

var ttl time.Duration
var a, b string

func init() {
	ttl = 24 * time.Hour
	a = "a"
	b = "b"
}

func TestPutWithNegativeTTLFails(t *testing.T) {
	c := makeCache(t, 1)

	negativeTTL := -1 * time.Hour
	err := c.Put(a, a, negativeTTL)
	assert.NotNil(t, err)
}

func TestPutEvictsOldEntries(t *testing.T) {
	c := makeCache(t, 1)

	err := c.Put(a, a, ttl)
	assert.Nil(t, err)

	v, ok := c.Get(a)
	assert.True(t, ok)
	assert.Equal(t, a, v)

	err = c.Put(b, b, ttl)
	assert.Nil(t, err)

	v, ok = c.Get(a)
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestGetRetrievesCorrectValue(t *testing.T) {
	c := makeCache(t, 2)

	c.Put(a, a, ttl)
	c.Put(b, b, ttl)

	v, ok := c.Get(a)
	assert.True(t, ok)
	assert.Equal(t, a, v)

	v, ok = c.Get(b)
	assert.True(t, ok)
	assert.Equal(t, b, v)
}

func TestGetWithMissingKeyFails(t *testing.T) {
	c := makeCache(t, 1)

	c.Put(a, a, ttl)

	v, ok := c.Get(a)
	assert.True(t, ok)
	assert.Equal(t, a, v)

	v, ok = c.Get(b)
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestGetWithExpiredEntryFails(t *testing.T) {
	c := makeCache(t, 1)

	c.Put(a, a, 50*time.Millisecond)

	time.Sleep(51 * time.Millisecond)

	v, ok := c.Get(a)
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestDelete(t *testing.T) {
	c := makeCache(t, 1)

	c.Put(a, a, ttl)

	v, ok := c.Get(a)
	assert.True(t, ok)
	assert.Equal(t, a, v)

	c.Delete(a)

	v, ok = c.Get(a)
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestPurge(t *testing.T) {
	c := makeCache(t, 2)

	c.Put(a, a, ttl)
	c.Put(b, b, ttl)

	v, ok := c.Get(a)
	assert.True(t, ok)
	assert.Equal(t, a, v)

	v, ok = c.Get(b)
	assert.True(t, ok)
	assert.Equal(t, b, v)

	c.Purge()

	v, ok = c.Get(a)
	assert.False(t, ok)
	assert.Equal(t, "", v)

	v, ok = c.Get(b)
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func makeCache(t *testing.T, size int) cache.Cache {
	c, err := NewCache(size)
	assert.Nil(t, err)
	return c
}
