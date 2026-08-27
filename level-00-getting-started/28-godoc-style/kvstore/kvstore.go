package kvstore

// Store is a minimal in-memory key-value store.
//
// The zero value is NOT ready to use — construct one with [New].
type Store struct {
	data map[string]string
}

// New returns a ready-to-use, empty Store.
func New() *Store {
	return &Store{data: make(map[string]string)}
}

// Set stores value under key, overwriting any existing value.
func (s *Store) Set(key, value string) {
	s.data[key] = value
}

// Get returns the value stored under key, and whether it was present.
func (s *Store) Get(key string) (value string, ok bool) {
	value, ok = s.data[key]
	return value, ok
}

// Delete removes key from the store, if present. Deleting a key that
// doesn't exist is a no-op, not an error.
func (s *Store) Delete(key string) {
	delete(s.data, key)
}

// Len returns the number of keys currently stored.
func (s *Store) Len() int {
	return len(s.data)
}
