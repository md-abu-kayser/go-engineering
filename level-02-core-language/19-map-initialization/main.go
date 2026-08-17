package main

import "fmt"

func main() {
	counts := map[string]int{"MapInitialization": 1}
	counts["total"]++
	fmt.Printf("Map Initialization: %v\n", counts)
}
