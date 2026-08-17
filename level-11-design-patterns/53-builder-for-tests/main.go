package main

import "fmt"

func main() {
	input := "Builder For Tests"
	result := len(input) > 0
	fmt.Printf("subject=%q passes-basic-check=%t\n", input, result)
}
