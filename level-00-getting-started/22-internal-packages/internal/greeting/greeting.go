// Package greeting is an INTERNAL package.
//
// Because it lives under a directory literally named "internal", the Go
// compiler only allows it to be imported by code rooted at the parent of
// that "internal" directory — in this case, anything under
// level-00-getting-started/22-internal-packages/...
//
// This is enforced by the toolchain itself, not just a naming convention.
package greeting

// Hello returns a greeting for name.
func Hello(name string) string {
	return "Hello from an internal package, " + name + "!"
}
