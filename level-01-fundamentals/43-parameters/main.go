// Lesson 43: Parameters
//
// Goal: Understand that Go is ALWAYS pass-by-value — including for
// slices and maps, which only APPEAR to be mutable through function
// calls because what's copied is a small "header" pointing at shared
// underlying data.
package main

import "fmt"

// tryToDouble receives a COPY of n — modifying it here has ZERO effect
// on whatever variable the caller passed in.
func tryToDouble(n int) {
	n *= 2
	fmt.Printf("  inside tryToDouble: n is now %d\n", n)
}

// modifyFirstElement receives a COPY of the slice's HEADER (pointer,
// length, capacity) — but that header's pointer points at the SAME
// underlying array the caller has. Modifying an ELEMENT through it
// really does affect the caller's data.
func modifyFirstElement(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
}

// appendToSlice shows the OTHER half of the story: append can allocate
// a completely NEW underlying array if the old one runs out of
// capacity — and the caller's original slice header has no way to know
// that happened, since ITS copy of the header is unaffected.
func appendToSlice(s []int) {
	s = append(s, 100) // may or may not affect the caller's slice — see README
}

// modifyMapEntry shows a map behaves like a slice's header here: the
// map "value" is really a small pointer to shared underlying storage,
// so writes through it ARE visible to the caller.
func modifyMapEntry(m map[string]int) {
	m["new"] = 1
}

func main() {
	fmt.Println("=== Parameters ===")
	fmt.Println("----------------------------------")

	fmt.Println("--- Plain values: always copied, never affected ---")
	x := 5
	fmt.Printf("before tryToDouble: x = %d\n", x)
	tryToDouble(x)
	fmt.Printf("after tryToDouble:  x = %d (COMPLETELY unchanged)\n", x)

	fmt.Println("\n--- Slices: the HEADER is copied, but it points at SHARED data ---")
	nums := []int{1, 2, 3}
	fmt.Printf("before modifyFirstElement: nums = %v\n", nums)
	modifyFirstElement(nums)
	fmt.Printf("after modifyFirstElement:  nums = %v (element WAS changed — shared array)\n", nums)

	fmt.Println("\n--- append() inside a function may NOT affect the caller ---")
	small := make([]int, 2, 2) // length 2, capacity 2 — NO room to grow in place
	small[0], small[1] = 1, 2
	fmt.Printf("before appendToSlice: small = %v (len=%d, cap=%d)\n", small, len(small), cap(small))
	appendToSlice(small)
	fmt.Printf("after appendToSlice:  small = %v (UNCHANGED — append had to allocate a new array)\n", small)

	fmt.Println("\n--- Maps: behave like slices' shared-data property, but for the WHOLE map ---")
	ages := map[string]int{"Alice": 30}
	fmt.Printf("before modifyMapEntry: ages = %v\n", ages)
	modifyMapEntry(ages)
	fmt.Printf("after modifyMapEntry:  ages = %v (new entry IS visible — shared storage)\n", ages)
}
