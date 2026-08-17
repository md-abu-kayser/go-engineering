package main

import "fmt"

func summarizeDomainPackages() (string, int) {
	topic := "Domain Packages"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDomainPackages()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
