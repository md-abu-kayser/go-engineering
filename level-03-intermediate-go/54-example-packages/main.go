package main

import "fmt"

func ExamplePackages() string {
	const topic = "Example Packages"
	return topic
}

func main() {
	fmt.Println(ExamplePackages())
}
