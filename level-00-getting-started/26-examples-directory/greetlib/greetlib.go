// Package greetlib is a tiny, standalone library — the kind of thing
// this lesson's examples/ directory exists to demonstrate usage of.
package greetlib

// Greeting represents a configurable greeter.
type Greeting struct {
	Prefix string // e.g. "Hello", "Hi", "Welcome"
}

// For returns the greeting for name, using g's Prefix.
// An empty Prefix defaults to "Hello".
func (g Greeting) For(name string) string {
	prefix := g.Prefix
	if prefix == "" {
		prefix = "Hello"
	}
	return prefix + ", " + name + "!"
}
