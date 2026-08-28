// Lesson 11: Integers
//
// Goal: Understand Go's signed integer types, the platform-dependent
// size of plain `int`, integer division/modulo, and overflow behavior.
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("=== Integers ===")
	fmt.Println("----------------------------------")

	// Plain `int` is the type you'll use almost everywhere — its SIZE
	// (32 or 64 bits) depends on the target platform. unsafe.Sizeof
	// reports it directly, in bytes, rather than us assuming.
	var i int = 42
	fmt.Printf("int          : %d (size on this machine: %d bits)\n", i, 8*unsafe.Sizeof(i))

	// The explicitly-sized variants — use these when the EXACT size
	// genuinely matters (binary file formats, network protocols,
	// matching a C struct layout, etc.), not by default.
	var i8 int8 = 127
	var i16 int16 = 32000
	var i32 int32 = 2000000000
	var i64 int64 = 9000000000000000000
	fmt.Printf("int8         : %d (range: %d to %d)\n", i8, math.MinInt8, math.MaxInt8)
	fmt.Printf("int16        : %d (range: %d to %d)\n", i16, math.MinInt16, math.MaxInt16)
	fmt.Printf("int32        : %d (range: %d to %d)\n", i32, math.MinInt32, math.MaxInt32)
	fmt.Printf("int64        : %d (range: %d to %d)\n", i64, int64(math.MinInt64), int64(math.MaxInt64))

	// Integer division TRUNCATES toward zero — it does NOT round.
	fmt.Println("\n--- Integer division & modulo ---")
	fmt.Printf("7 / 2   = %d (truncated, not rounded)\n", 7/2)
	fmt.Printf("7 %% 2   = %d (remainder)\n", 7%2)
	fmt.Printf("-7 / 2  = %d (truncates TOWARD ZERO, not toward negative infinity)\n", -7/2)
	fmt.Printf("-7 %% 2  = %d\n", -7%2)

	// Overflow WRAPS AROUND silently — Go does not panic or error on
	// integer overflow by default.
	fmt.Println("\n--- Overflow ---")
	var small int8 = 127 // int8's max value
	small++
	fmt.Printf("int8(127) + 1 = %d (wrapped around to the MINIMUM value, silently)\n", small)
}
