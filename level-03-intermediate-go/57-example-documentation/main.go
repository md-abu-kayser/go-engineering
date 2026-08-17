package main

import "fmt"

func ExampleDocumentation() string {
	const topic = "Example Documentation"
	return topic
}

func main() {
	fmt.Println(ExampleDocumentation())
}
