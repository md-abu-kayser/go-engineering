// Lesson 03: Zero Values
//
// Goal: See the zero value for every major kind of type Go has — proof
// that "declared but not explicitly initialized" NEVER means garbage or
// undefined memory in Go, unlike some other languages.
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var r rune
	var by byte

	var slice []int
	var m map[string]int
	var p *int
	var ch chan int
	var fn func()
	var iface interface{}

	type point struct{ X, Y int }
	var pt point

	fmt.Println("=== Zero Values ===")
	fmt.Println("----------------------------------")
	fmt.Printf("int              : %v\n", i)
	fmt.Printf("float64          : %v\n", f)
	fmt.Printf("bool             : %v\n", b)
	fmt.Printf("string           : %q (empty, not nil)\n", s)
	fmt.Printf("rune             : %v\n", r)
	fmt.Printf("byte             : %v\n", by)
	fmt.Printf("[]int (slice)    : %v (nil: %t, len: %d)\n", slice, slice == nil, len(slice))
	fmt.Printf("map[string]int   : %v (nil: %t)\n", m, m == nil)
	fmt.Printf("*int (pointer)   : %v (nil: %t)\n", p, p == nil)
	fmt.Printf("chan int         : %v (nil: %t)\n", ch, ch == nil)
	fmt.Printf("func()           : nil: %t\n", fn == nil)
	fmt.Printf("interface{}      : %v (nil: %t)\n", iface, iface == nil)
	fmt.Printf("struct{X,Y int}  : %+v (fields get THEIR zero values too)\n", pt)
}
