// Package main demonstrates Go doc comments — the plain comments directly
// above a declaration that `go doc` and pkg.go.dev turn into documentation.
//
// Lesson 13: go doc
//
// Goal: Write proper doc comments and read them back with the `go doc`
// command, the same way you'd look up documentation for any standard
// library package.
package main

import "fmt"

// Greeter produces greetings for a configured language.
//
// The zero value of Greeter is ready to use and defaults to English.
type Greeter struct {
	// Language is a two-letter language code, e.g. "en" or "bn".
	// An empty Language behaves as "en".
	Language string
}

// Greet returns a greeting for name, in the Greeter's configured language.
//
// Currently supported languages are "en" (English) and "bn" (Bangla);
// any other value falls back to English.
func (g Greeter) Greet(name string) string {
	switch g.Language {
	case "bn":
		return "হ্যালো, " + name + "!"
	default:
		return "Hello, " + name + "!"
	}
}

func main() {
	en := Greeter{}
	bn := Greeter{Language: "bn"}

	fmt.Println(en.Greet("Gopher"))
	fmt.Println(bn.Greet("Gopher"))

	fmt.Println("\nRun `go doc .` and `go doc Greeter.Greet` in this folder to see these")
	fmt.Println("comments rendered as documentation.")
}
