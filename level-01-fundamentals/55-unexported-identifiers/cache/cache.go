// Package cache demonstrates a REALISTIC use of unexported
// identifiers: hiding implementation details (a map and a mutex-like
// counter) behind a small, exported API — so callers can only interact
// with the cache through deliberately-designed methods, never by
// reaching into its internals directly.
package cache

// Cache is EXPORTED — the type itself is usable from other packages.
type Cache struct {
	// data is UNEXPORTED — even though Cache itself is exported,
	// its FIELD is not. Callers outside this package cannot read or
	// write cache.data directly, even though they can hold a *Cache.
	data map[string]string

	// hits is ALSO unexported — purely an internal bookkeeping detail
	// callers have no business touching directly.
	hits int
}

// New is the EXPORTED constructor — the only way for another package
// to get a properly-initialized *Cache, since `data` (unexported)
// can't be set directly from outside.
func New() *Cache {
	return &Cache{data: make(map[string]string)}
}

// Set is part of the small, deliberate EXPORTED API.
func (c *Cache) Set(key, value string) {
	c.data[key] = value
}

// Get is the other half of the exported API — and it's ALSO where the
// unexported `hits` field gets updated, entirely hidden from callers.
func (c *Cache) Get(key string) (string, bool) {
	c.hits++ // an internal detail callers never see or control directly
	value, ok := c.data[key]
	return value, ok
}

// Hits is a DELIBERATE, controlled window into the otherwise-hidden
// hits field — exported specifically as a read-only accessor, rather
// than exposing the field itself.
func (c *Cache) Hits() int {
	return c.hits
}
