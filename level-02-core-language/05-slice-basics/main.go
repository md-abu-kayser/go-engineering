package main

import "fmt"

func main() {
	values := make([]int, 0, 4)
	for i := 1; i <= 4; i++ {
		values = append(values, i*i)
	}
	fmt.Printf("Slice Basics: %v len=%d cap=%d\n", values, len(values), cap(values))
}
