package main

import "fmt"

func main() {
	counts := map[string]int{"MapSharding": 1}
	counts["total"]++
	fmt.Printf("Map Sharding: %v\n", counts)
}
