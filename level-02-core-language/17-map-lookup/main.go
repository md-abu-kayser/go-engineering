package main

import "fmt"

func main() {
	counts := map[string]int{"MapLookup": 1}
	counts["total"]++
	fmt.Printf("Map Lookup: %v\n", counts)
}
