// Lesson 06: iota
//
// Goal: Use `iota` to build auto-incrementing constant enumerations —
// including the classic bit-shifted flags pattern — instead of manually
// numbering every constant by hand.
package main

import "fmt"

// A basic iota sequence: iota starts at 0 and increments by 1 for each
// line within the const block, so this is exactly Sunday=0, Monday=1, ...
type weekday int

const (
	Sunday weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// SKIPPING a value: the blank identifier consumes one iota "slot"
// without creating a named constant for it.
const (
	Bronze   = iota + 1 // 1 (iota starts at 0 again in this NEW const block)
	Silver              // 2
	_                   // 3 (skipped on purpose)
	Platinum            // 4
)

// The classic bit-flag pattern: shifting 1 left by iota gives each
// constant a distinct, independent BIT, so they can be combined with |
// and tested with &.
type permission uint8

const (
	PermRead    permission = 1 << iota // 1 << 0 = 1
	PermWrite                          // 1 << 1 = 2
	PermExecute                        // 1 << 2 = 4
)

func main() {
	fmt.Println("=== iota ===")
	fmt.Println("----------------------------------")
	fmt.Printf("Sunday=%d Monday=%d ... Saturday=%d\n", Sunday, Monday, Saturday)
	fmt.Printf("Bronze=%d Silver=%d Platinum=%d (value 3 skipped)\n", Bronze, Silver, Platinum)

	fmt.Printf("PermRead=%d PermWrite=%d PermExecute=%d\n", PermRead, PermWrite, PermExecute)

	// Combine flags with | (bitwise OR) — this is THE reason the
	// bit-shift pattern exists: distinct bits can be combined into one
	// value that represents multiple flags at once.
	readWrite := PermRead | PermWrite
	fmt.Printf("readWrite (Read|Write) = %d\n", readWrite)

	// Test for a specific flag with & (bitwise AND).
	fmt.Printf("readWrite has PermWrite?   %t\n", readWrite&PermWrite != 0)
	fmt.Printf("readWrite has PermExecute? %t\n", readWrite&PermExecute != 0)
}
