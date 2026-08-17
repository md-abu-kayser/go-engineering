package main

import "fmt"

func summarizeLoggingPackages() (string, int) {
	topic := "Logging Packages"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeLoggingPackages()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
