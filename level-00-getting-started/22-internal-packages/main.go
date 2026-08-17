package main

import "fmt"

func InternalPackages() string {
	const topic = "Internal Packages"
	return topic
}

func main() {
	fmt.Println(InternalPackages())
}
