// Lesson 33: range Loop
//
// Goal: Use `for range` over every common collection type — slice,
// array, string, and map — and know exactly what each one yields.
package main

import "fmt"

func main() {
	fmt.Println("=== range Loop ===")
	fmt.Println("----------------------------------")

	// Slice: range yields (index, value) pairs, in order.
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("--- range over a slice ---")
	for i, fruit := range fruits {
		fmt.Printf("  [%d] = %s\n", i, fruit)
	}

	// Only the index is needed sometimes — just omit the value.
	fmt.Println("\n--- range with index only ---")
	for i := range fruits {
		fmt.Printf("  index %d\n", i)
	}

	// Only the value is needed more often — use the blank identifier
	// to discard the index explicitly.
	fmt.Println("\n--- range with value only (index discarded) ---")
	for _, fruit := range fruits {
		fmt.Printf("  %s\n", fruit)
	}

	// Array: works identically to a slice.
	fmt.Println("\n--- range over an array ---")
	var nums [3]int = [3]int{10, 20, 30}
	for i, n := range nums {
		fmt.Printf("  [%d] = %d\n", i, n)
	}

	// String: range decodes UTF-8 and yields (byte index, rune) pairs —
	// exactly the behavior lesson 09 already covered in depth.
	fmt.Println("\n--- range over a string (decodes runes) ---")
	for i, r := range "Hi, 世" {
		fmt.Printf("  byte %d: %c\n", i, r)
	}

	// Map: range yields (key, value) pairs — but in NO GUARANTEED
	// ORDER. Go deliberately randomizes map iteration order across
	// runs, precisely so nobody accidentally relies on an order that
	// was never promised.
	fmt.Println("\n--- range over a map (order NOT guaranteed) ---")
	ages := map[string]int{"Alice": 30, "Bob": 25}
	count := 0
	for name, age := range ages {
		fmt.Printf("  %s is %d\n", name, age)
		count++
	}
	fmt.Printf("  (%d entries total, in unspecified order)\n", count)
}
