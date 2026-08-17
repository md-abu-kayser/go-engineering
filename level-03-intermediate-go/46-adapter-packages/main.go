package main

import "fmt"

func summarizeAdapterPackages() (string, int) {
	topic := "Adapter Packages"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAdapterPackages()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
