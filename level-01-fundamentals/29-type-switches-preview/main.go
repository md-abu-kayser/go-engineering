// Lesson 29: Type Switches (Preview)
//
// Goal: Get a first, working look at Go's TYPE switch — a special
// switch form that branches on a value's DYNAMIC TYPE rather than its
// value. Interfaces themselves are a later topic; this lesson only
// covers the syntax, using Go's built-in `any` (interface{}) type.
package main

import "fmt"

// describe accepts ANY value at all — `any` is Go's built-in alias for
// the empty interface{}, meaning "a value of literally any type." A
// full explanation of interfaces comes in a later level; for now,
// think of `any` as "the type of parameter this lesson needs to accept
// values of every different kind."
func describe(v any) string {
	// switch v := v.(type) is the TYPE SWITCH syntax: v.(type) can
	// ONLY appear directly inside a switch like this. Each case
	// matches a specific concrete TYPE, not a value.
	switch v := v.(type) {
	case int:
		return fmt.Sprintf("an int: %d (doubled: %d)", v, v*2)
	case string:
		return fmt.Sprintf("a string: %q (length: %d)", v, len(v))
	case bool:
		return fmt.Sprintf("a bool: %t", v)
	case nil:
		return "a nil value"
	default:
		return fmt.Sprintf("something else entirely: %v (%T)", v, v)
	}
}

func main() {
	fmt.Println("=== Type Switches (Preview) ===")
	fmt.Println("----------------------------------")

	values := []any{42, "Gopher", true, 3.14, nil}
	for _, v := range values {
		fmt.Println(describe(v))
	}
}
