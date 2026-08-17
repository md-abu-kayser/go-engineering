package main

import "fmt"

func PackageInitializationOrder() string {
	const topic = "Package Initialization Order"
	return topic
}

func main() {
	fmt.Println(PackageInitializationOrder())
}
