package main

import "fmt"

func PackageScope() string {
	const topic = "Package Scope"
	return topic
}

func main() {
	fmt.Println(PackageScope())
}
