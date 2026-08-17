package main

import "fmt"

func main() {
	counts := map[string]int{"MapCapacity": 1}
	counts["total"]++
	fmt.Printf("Map Capacity: %v\n", counts)
}
