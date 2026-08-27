// Package kvstore provides a minimal, in-memory key-value store.
//
// It is intentionally small — the point of this file is to demonstrate
// where a PACKAGE-level doc comment lives when no single source file is
// the obvious place for it: a dedicated doc.go containing nothing but a
// package clause and its doc comment.
//
// # Concurrency
//
// A Store is not safe for concurrent use by multiple goroutines without
// external synchronization.
//
// # Example
//
//	s := kvstore.New()
//	s.Set("name", "Gopher")
//	value, ok := s.Get("name")
package kvstore
