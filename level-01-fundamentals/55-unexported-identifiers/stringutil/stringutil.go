// Package stringutil demonstrates UNEXPORTED identifiers used
// deliberately for encapsulation: internal helper logic that outside
// code should never depend on directly.
package stringutil

import "strings"

// normalize is UNEXPORTED (lowercase) — an internal implementation
// detail of Slugify below. Outside packages have NO way to call this
// directly, which is exactly the point: it can be renamed, rewritten,
// or removed entirely without ever affecting anyone importing this
// package, since it was never part of the public contract to begin with.
func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// Slugify is the EXPORTED entry point — the only thing outside code
// can actually call. It uses normalize internally, but callers never
// need to know that detail exists at all.
func Slugify(title string) string {
	cleaned := normalize(title)
	return strings.ReplaceAll(cleaned, " ", "-")
}

// cache is an UNEXPORTED package-level variable — genuinely private
// state that only code inside THIS package can read or modify, no
// matter what package imports it.
var cache = map[string]string{}

// SlugifyCached wraps Slugify with a simple cache — cache itself
// remains completely invisible and untouchable from outside, which is
// exactly what makes it safe to add here without it becoming part of
// this package's public API surface.
func SlugifyCached(title string) string {
	if cached, ok := cache[title]; ok {
		return cached
	}
	result := Slugify(title)
	cache[title] = result
	return result
}
