package main

import "fmt"

func main() {
	counts := map[string]int{"MapRuntimeAwareness": 1}
	counts["total"]++
	fmt.Printf("Map Runtime Awareness: %v\n", counts)
}
