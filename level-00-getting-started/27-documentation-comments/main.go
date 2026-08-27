// Package main demonstrates the Go doc comment FORMATTING rules
// introduced with the "gofmt-aware" doc comment renderer: headings,
// lists, links, and indented code blocks — all inside plain comments.
//
// Lesson 27: Documentation Comments
//
// Goal: Learn the specific syntax Go's doc comment renderer understands,
// beyond just "write a sentence above a declaration".
package main

import "fmt"

// Temperature represents a temperature reading.
//
// # Units
//
// Temperature always stores its value in Celsius internally. Use
// [Temperature.Fahrenheit] to convert for display when needed.
//
// # Example
//
// A typical reading is created like this:
//
//	t := Temperature{Celsius: 21.5}
//	fmt.Println(t.Fahrenheit())
//
// See also the Go blog post on doc comments: https://go.dev/doc/comment
type Temperature struct {
	Celsius float64
}

// Fahrenheit converts t to degrees Fahrenheit.
//
// Deprecated: prefer working in Celsius directly; this method is kept
// only for display purposes in US-locale output.
func (t Temperature) Fahrenheit() float64 {
	return t.Celsius*9/5 + 32
}

func main() {
	t := Temperature{Celsius: 21.5}
	fmt.Printf("%.1f°C = %.1f°F\n", t.Celsius, t.Fahrenheit())
	fmt.Println("\nRun `go doc -all .` in this folder to see how the comments above render.")
}
