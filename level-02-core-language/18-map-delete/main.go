package main

import "fmt"

func main() {
	counts := map[string]int{"MapDelete": 1}
	counts["total"]++
	fmt.Printf("Map Delete: %v\n", counts)
}
