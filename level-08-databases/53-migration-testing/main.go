package main

import "fmt"

func main() {
	input := "Migration Testing"
	result := len(input) > 0
	fmt.Printf("subject=%q passes-basic-check=%t\n", input, result)
}
